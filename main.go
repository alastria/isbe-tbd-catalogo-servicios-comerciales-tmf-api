package main

import (
	"context"
	"flag" // Added
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"

	"log/slog"

	"github.com/cloudflare/tableflip"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/hesusruiz/isbetmf/internal/errl"
	"github.com/hesusruiz/isbetmf/internal/sqlogger"
	"github.com/hesusruiz/isbetmf/pdp"
	fiberhandler "github.com/hesusruiz/isbetmf/tmfserver/handler/fiber"
	repository "github.com/hesusruiz/isbetmf/tmfserver/repository"
	service "github.com/hesusruiz/isbetmf/tmfserver/service"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	var debugFlag bool
	var verifierServer string
	var remoteTMFServer string
	var proxyEnabled bool
	var restartHour, restartMinute int

	// Parse command-line flags
	flag.BoolVar(&debugFlag, "d", true, "Enable debug logging")

	flag.StringVar(&verifierServer, "verifier", "", "Full URL of the verifier which signs access tokens")
	flag.StringVar(&remoteTMFServer, "remote", "", "Full URL of the remote TMForum server to proxy requests to")
	flag.BoolVar(&proxyEnabled, "proxy", false, "Enable proxy functionality")
	flag.IntVar(&restartHour, "rh", 3, "Restart program every day at this hour")
	flag.IntVar(&restartMinute, "rm", 0, "Restart program every day at this minute")
	flag.Parse()

	// ******************************************************
	// ******************************************************
	// Detect if we are running as PID=1 (most probably as init process in a container),
	// and act accordingly.
	ourPid := os.Getpid()

	// Get the name fo our executable
	ourExecPath, err := os.Executable()
	if err != nil {
		panic(err)
	}

	args := os.Args[1:]

	runAsInit := false
	if ourPid == 1 {
		runAsInit = true
	} else {
		if len(args) > 0 && args[0] == "init" {
			runAsInit = true
			args = args[1:]
		}
	}

	if runAsInit {
		// We are the init process in a container, or testing the functionality
		fmt.Println("We are the init process! PID:", ourPid)

		// Pass to child all arguments following the "init" entry (first argument after program name)
		cmd := exec.Command(ourExecPath, args...)

		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		// Set ProcessGroupID for child process as init process. Both will be under same process group
		cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}

		// We need notification of all relevant signals
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

		done := make(chan bool, 1)

		// Forward all signals to the child in a goroutine
		go func() {
			for sig := range sigs {

				if sig == syscall.SIGHUP {
					fmt.Println("init process received SIGHUP")
				} else {
					fmt.Printf("init process received %v signal\n", sig)
				}

				if cmd.Process != nil {
					fmt.Printf("received %v signal for PID %v\n", sig, cmd.Process.Pid)
					_ = cmd.Process.Signal(sig)
				}

				if sig == syscall.SIGTERM || sig == syscall.SIGINT {
					// We are done, and the init process will terminate
					fmt.Println("using DONE channel to terminate init process")
					done <- true
				}
			}
		}()

		// Enter in a goroutine an infinite loop reaping the zombie children
		go func() {
			for {
				var ws syscall.WaitStatus
				pid, err := syscall.Wait4(-1, &ws, syscall.WNOHANG, nil)
				if pid <= 0 || err != nil {
					time.Sleep(1 * time.Second)
				} else {
					fmt.Printf("reaped zombie %v\n", pid)
				}
			}

		}()

		// Start the child (a fork of ourselves) without waiting for termination
		fmt.Println("Starting child process")
		if err := cmd.Start(); err != nil {
			log.Fatalf("Failed to start child process: %v", err)
		}

		fmt.Println("awaiting signal to terminate init process")
		<-done
		fmt.Println("exiting init process")
		return

	}
	// ******************************************************
	// ******************************************************
	// ******************************************************

	// Configure the slog logger
	var logLevel slog.Level
	if debugFlag {
		logLevel = slog.LevelDebug
	} else {
		logLevel = slog.LevelInfo
	}

	// Initialize the custom SQLogHandler
	logOptions := &sqlogger.Options{
		Level: &logLevel,
	}
	// Check if the logs should be colored
	if os.Getenv("ISBETMF_LOGS_NOCOLOR") == "true" {
		logOptions.NoColor = true
	}

	sqlog, err := sqlogger.NewSQLogHandler(logOptions)
	if err != nil {
		slog.Error("failed to initialize SQLogHandler", slog.Any("error", err))
		os.Exit(1)
	}
	defer sqlog.Close()
	slog.SetDefault(slog.New(sqlog))

	pid := os.Getpid()
	slog.Info("Process started", "PID", pid)

	// Get the url of the verifier from command line (priority) or environment variable
	if verifierServer == "" {
		verifierServer = os.Getenv("ISBETMF_VERIFIER")
		if verifierServer == "" {
			verifierServer = "https://verifier.dome-marketplace.eu"
		}
	}
	slog.Info("Verifier server", slog.String("url", verifierServer))

	if remoteTMFServer == "" {
		remoteTMFServer = os.Getenv("ISBETMF_REMOTE_SERVER")
		if remoteTMFServer == "" {
			remoteTMFServer = "https://tmf.dome-marketplace-sbx.eu"
		}
	}
	slog.Info("Remote TMF server", slog.String("url", remoteTMFServer))

	// Get the proxyEnabled from command line (priority) or environment variable
	if !proxyEnabled { // Only check env if not set by flag
		if os.Getenv("ISBETMF_PROXY_ENABLED") == "true" {
			proxyEnabled = true
		}
	}
	slog.Info("Proxy", slog.Bool("enabled", proxyEnabled))

	// Connect to the database
	db, err := sqlx.Connect("sqlite3", "data/isbetmf.db")
	if err != nil {
		slog.Error("failed to connect to database", slog.Any("error", err))
		os.Exit(1)
	}
	defer db.Close()

	slog.Info("About to create tables if they do not exist")
	err = repository.CreateTables(db)
	if err != nil {
		slog.Error("failed to create tables", slog.Any("error", err))
		os.Exit(1)
	}

	// Create the PDP (aka rules engine)
	rulesEngine, err := pdp.NewPDP(&pdp.Config{
		PolicyFileName: "auth_policies.star",
		Debug:          debugFlag,
	})
	if err != nil {
		slog.Error("failed to create rules engine", slog.Any("error", err))
		os.Exit(1)
	}

	// Create the service
	s := service.NewService(db, rulesEngine, verifierServer, proxyEnabled, remoteTMFServer)

	// Create Fiber app with custom configuration
	app := fiber.New(fiber.Config{
		AppName:      "ISBE-TMF API Server",
		ServerHeader: "ISBE-TMF",
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			slog.Error("Fiber error", slog.Any("error", err), slog.Int("status", code))
			return c.Status(code).JSON(fiber.Map{
				"error": err.Error(),
			})
		},
	})

	// Add middleware in order (order matters!)

	// 1. Recovery middleware - should be first to catch panics
	app.Use(recover.New(recover.Config{
		EnableStackTrace: debugFlag,
	}))

	// 2. Request ID middleware - for tracing requests
	app.Use(requestid.New(requestid.Config{
		Header: "X-Request-Id",
		Generator: func() string {
			return "req_" + time.Now().Format("20060102150405") + "_" + os.Getenv("HOSTNAME")
		},
	}))

	// 3. CORS middleware - enable cross-origin requests
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*", // Allow all origins as requested
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
		AllowHeaders:     "Origin,Content-Type,Accept,Authorization,X-Request-Id",
		AllowCredentials: false, // Set to false when AllowOrigins is "*"
		ExposeHeaders:    "X-Request-Id",
		MaxAge:           86400, // 24 hours
	}))

	// 4. Compression middleware - compress responses
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))

	// // 5. Rate limiting middleware - prevent abuse
	// app.Use(limiter.New(limiter.Config{
	// 	Max:        10000000,        // Maximum number of requests
	// 	Expiration: 1 * time.Minute, // Time window
	// 	KeyGenerator: func(c *fiber.Ctx) string {
	// 		// Use client IP for rate limiting
	// 		return c.IP()
	// 	},
	// 	LimitReached: func(c *fiber.Ctx) error {
	// 		return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
	// 			"error": "Rate limit exceeded",
	// 		})
	// 	},
	// }))

	// 6. Logger middleware - log requests and replies
	app.Use(sqlogger.FiberRequestLogger)

	// Note: Timeout middleware removed due to API changes in Fiber v2.52.9
	// Consider implementing request timeouts at the handler level if needed

	// Serve the OpenAPI UI
	app.Static("/oapiv5", "./www/oapiv5")
	app.Static("/oapiv4", "./www/oapiv4")

	// Create handler and set the routes for the APIs
	h := fiberhandler.NewHandler(s)
	h.RegisterRoutes(app)

	// TABLEFLIP for seamless restarts and upgrades
	upg, _ := tableflip.New(tableflip.Options{
		PIDFile: "isbetmf.pid",
	})
	defer upg.Stop()

	// Listen for the process signal to trigger the tableflip upgrade.
	go func() {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, syscall.SIGHUP)
		for range sig {
			fmt.Println("Received SIGHUP, upgrading...")
			upg.Upgrade()
		}
	}()

	// Schedule restarts/upgrades every night, the exact time does not matter
	targetHour := restartHour
	targetMinute := restartMinute
	targetSecond := 0

	go func() {
		for {
			now := time.Now()

			// Calculate the next scheduled time
			nextRun := time.Date(
				now.Year(), now.Month(), now.Day(),
				targetHour, targetMinute, targetSecond, 0, now.Location(),
			)

			// If the next run time is in the past, schedule it for the next day
			if nextRun.Before(now) {
				nextRun = nextRun.Add(24 * time.Hour)
			}

			fmt.Printf("Time now: %v\n", now)

			// Calculate the duration until the next run
			duration := time.Until(nextRun)
			fmt.Printf("Next restart/upgrade scheduled at: %v\n", nextRun)

			// Wait until the next run time
			time.Sleep(duration)

			// Execute the function
			fmt.Println("Executing scheduled restart...")
			upg.Upgrade()
		}
	}()

	// Listen must be called before Ready
	ln, err := upg.Listen("tcp", "0.0.0.0:9991")
	if err != nil {
		panic(errl.Error(err))
	}
	defer ln.Close()

	// Start the server in a separate goroutine
	go func() {
		slog.Info("TMF API server starting", slog.String("port", ":9991"))
		err := app.Listener(ln)
		if err != nil {
			slog.Error("Error starting TMF API server", "error", errl.Error(err))
			panic(err)
		}
	}()

	// tableflip ready
	slog.Info("Tableflip ready")
	if err := upg.Ready(); err != nil {
		panic(errl.Error(err))
	}

	<-upg.Exit()

	// Make sure to set a deadline on exiting the process
	// after upg.Exit() is closed. No new upgrades can be
	// performed if the parent doesn't exit.
	time.AfterFunc(30*time.Second, func() {
		log.Println("Graceful shutdown timed out")
		os.Exit(1)
	})

	// Wait for connections to drain.
	err = app.ShutdownWithContext(context.Background())
	if err != nil {
		fmt.Println("Exiting with error:", errl.Error(err).Error())
	} else {
		fmt.Println("Exiting without error")
	}

}
