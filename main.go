package main

import (
	"flag" // Added
	"os"
	"time"

	"log/slog"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/hesusruiz/isbetmf/internal/sqlogger"
	"github.com/hesusruiz/isbetmf/pdp"
	fiberhandler "github.com/hesusruiz/isbetmf/tmfserver/handler/fiber"
	repository "github.com/hesusruiz/isbetmf/tmfserver/repository"
	service "github.com/hesusruiz/isbetmf/tmfserver/service"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Configure slog logger
	var debugFlag bool
	var verifierServer string
	var remoteTMFServer string
	var proxyEnabled bool

	flag.BoolVar(&debugFlag, "d", true, "Enable debug logging")
	flag.StringVar(&verifierServer, "verifier", "", "Full URL of the verifier which signs access tokens")
	flag.StringVar(&remoteTMFServer, "remote", "", "Full URL of the remote TMForum server to proxy requests to")
	flag.BoolVar(&proxyEnabled, "proxy", false, "Enable proxy functionality")
	flag.Parse()

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
	sqlog, err := sqlogger.NewSQLogHandler(logOptions)
	if err != nil {
		slog.Error("failed to initialize SQLogHandler", slog.Any("error", err))
		os.Exit(1)
	}
	defer sqlog.Close()
	slog.SetDefault(slog.New(sqlog))

	// slogHandler := slogor.NewHandler(os.Stdout, slogor.ShowSource(), slogor.SetLevel(logLevel))
	// slog.SetDefault(slog.New(slogHandler))

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

	// 5. Rate limiting middleware - prevent abuse
	app.Use(limiter.New(limiter.Config{
		Max:        100,             // Maximum number of requests
		Expiration: 1 * time.Minute, // Time window
		KeyGenerator: func(c *fiber.Ctx) string {
			// Use client IP for rate limiting
			return c.IP()
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"error": "Rate limit exceeded",
			})
		},
	}))

	// 6. Logger middleware - log requests
	app.Use(sqlogger.FiberRequestLogger)

	// Note: Timeout middleware removed due to API changes in Fiber v2.52.9
	// Consider implementing request timeouts at the handler level if needed

	// Serve the OpenAPI UI
	app.Static("/oapiv5", "./www/oapiv5")
	app.Static("/oapiv4", "./www/oapiv4")

	// Create handler and set the routes for the APIs
	h := fiberhandler.NewHandler(s)
	h.RegisterRoutes(app)

	// And start the server
	slog.Info("TMF API server starting", slog.String("port", ":9991"))
	app.Listen("0.0.0.0:9991")

}
