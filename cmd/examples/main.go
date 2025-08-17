package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/zkfmapf123/lambda-pods/configs"
	"github.com/zkfmapf123/lambda-pods/internal"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	PORT     = internal.GetProcessEnv()["PORT"]
	APP_NAME = internal.GetProcessEnv()["APP_NAME"]
	DB_HOST  = internal.GetProcessEnv()["DB_HOST"]
	DB_PORT  = internal.GetProcessEnv()["DB_PORT"]
	DB_USER  = internal.GetProcessEnv()["DB_USER"]
	DB_PASS  = internal.GetProcessEnv()["DB_PASS"]
	DB_NAME  = internal.GetProcessEnv()["DB_NAME"]
	ENV      = internal.GetProcessEnv()["ENV"]
)

func main() {
	app := fiber.New(fiber.Config{
		// Prefork 상태에서는 Gracefully shutdown이 처리되지 않는다.
		// os.Signal은 현재 프로세스(워커) 에서만 받아지고 ->  다른 워커들은 신호를 받지 못한다... (즉 프로세스간 통신 IPC 신호전파 메커니즘 없음)
		Prefork:       false,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "fiber",
		AppName:       APP_NAME,
	})

	// global settings
	logger := internal.NewLogger(ENV)
	db := configs.MustInitDB(DB_HOST, DB_PORT, DB_USER, DB_PASS, DB_NAME)

	setDefaultRouter(app)
	setMiddleware(app)
	setRouters(app, RouterParams{
		db:     db,
		logger: logger,
	})

	go func() {
		if err := app.Listen(":" + PORT); err != nil {
			logger.Error("Server error", zap.Error(err))
		}
	}()

	/////////////////////////////////////////////// Gracefully Shutdown ///////////////////////////////////////////////
	q := make(chan os.Signal, 1)
	signal.Notify(q, os.Interrupt, syscall.SIGTERM)

	<-q
	log.Println("Received shutdown signal...")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	log.Println("Shutting down server...")
	if err := app.ShutdownWithContext(ctx); err != nil {
		log.Printf("Failed to shutdown server: %v\n", err)
	} else {
		log.Println("Server shutdown complete")
	}

	// Database Close
	if err := configs.CloseDB(db); err != nil {
		log.Printf("Failed to close database: %v\n", err)
	}

	// Logger Close
	defer logger.Sync()

	log.Println("Server shutdown complete")
}

func setMiddleware(app *fiber.App) {
	app.Use(logger.New())
}

func setDefaultRouter(app *fiber.App) {
	app.
		Get("/ping", func(c *fiber.Ctx) error {
			return c.SendString("success")
		}).
		// liveness probe
		Get("/livenss", func(c *fiber.Ctx) error {
			return c.SendString("liveness")
		}).
		// readness probe
		Get("/readiness", func(c *fiber.Ctx) error {
			return c.SendString("readiness")
		})
}

// ////////////////////////////////////////////////////// Router ////////////////////////////////////////////////////////
type RouterParams struct {
	db     *gorm.DB
	logger *zap.Logger
}

func setRouters(app *fiber.App, params RouterParams) {
	// gracefully shutdown test
	// app.Get("/test", func(c *fiber.Ctx) error {
	// 	time.Sleep(15 * time.Second)
	// 	return c.SendString("test")
	// })

	app.Get("/test", func(c *fiber.Ctx) error {
		params.logger.Info("test", zap.String("username", "leedonggyu"), zap.Int("age", 94), zap.String("job", "devops"))
		return c.SendString("test")
	})
}
