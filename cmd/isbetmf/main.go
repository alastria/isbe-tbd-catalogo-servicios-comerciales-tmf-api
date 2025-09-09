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
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/hesusruiz/isbetmf/pdp"
	fiberhandler "github.com/hesusruiz/isbetmf/tmfserver/handler/fiber"
	repository "github.com/hesusruiz/isbetmf/tmfserver/repository"
	service "github.com/hesusruiz/isbetmf/tmfserver/service"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"gitlab.com/greyxor/slogor"
)

func main() {
	// Configure slog logger
	var debugFlag bool
	var verifierServer string
	flag.BoolVar(&debugFlag, "d", false, "Enable debug logging")
	flag.StringVar(&verifierServer, "verifier", "", "Full URL of the verifier which signs access tokens")
	flag.Parse()

	// Get the url of the verifier from command line (priority) or environment variable
	if verifierServer == "" {
		verifierServer = os.Getenv("ISBETMF_VERIFIER")
		if verifierServer == "" {
			verifierServer = "https://verifier.dome-marketplace.eu"
		}
	}

	slog.Info("Verifier server", slog.String("verifierServer", verifierServer))

	// Use debug level until production
	debugFlag = true

	var logLevel slog.Level
	if debugFlag {
		logLevel = slog.LevelDebug
	} else {
		logLevel = slog.LevelInfo
	}

	handler := slogor.NewHandler(os.Stdout, slogor.ShowSource(), slogor.SetLevel(logLevel))
	slog.SetDefault(slog.New(handler))

	// Connect to the database
	db, err := sqlx.Connect("sqlite3", "isbetmf.db")
	if err != nil {
		slog.Error("failed to connect to database", slog.Any("error", err))
		os.Exit(1)
	}
	defer db.Close()

	// Create the table if it doesn't exist
	_, err = db.Exec(repository.CreateTMFTableSQL)
	if err != nil {
		slog.Error("failed to create table", slog.Any("error", err))
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
	s := service.NewService(db, rulesEngine, verifierServer)

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
	app.Use(logger.New(logger.Config{
		Format:     "${time} ${status} - ${method} ${path} (${ip}) - ${latency}\n",
		TimeFormat: "2006-01-02 15:04:05",
		TimeZone:   "UTC",
	}))

	// Note: Timeout middleware removed due to API changes in Fiber v2.52.9
	// Consider implementing request timeouts at the handler level if needed

	// Serve the OpenAPI UI
	app.Static("/oapi", "./www/oapiui")

	// Create handler and set the routes for the APIs
	h := fiberhandler.NewHandler(s)
	h.RegisterRoutes(app)

	// And start the server
	slog.Info("TMF API server starting", slog.String("port", ":9991"))
	app.Listen("0.0.0.0:9991")

}
