package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/zkfmapf123/lambda-pods/configs"
	"github.com/zkfmapf123/lambda-pods/internal"
)

var (
	PORT     = internal.GetProcessEnv()["PORT"]
	APP_NAME = internal.GetProcessEnv()["APP_NAME"]
	DB_HOST  = internal.GetProcessEnv()["DB_HOST"]
	DB_PORT  = internal.GetProcessEnv()["DB_PORT"]
	DB_USER  = internal.GetProcessEnv()["DB_USER"]
	DB_PASS  = internal.GetProcessEnv()["DB_PASS"]
	DB_NAME  = internal.GetProcessEnv()["DB_NAME"]
)

func main() {
	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "fiber",
		AppName:       APP_NAME,
	})

	configs.MustInitDB(DB_HOST, DB_PORT, DB_USER, DB_PASS, DB_NAME)

	setDefaultRouter(app)
	setMiddleware(app)
	setRouters(app)

	go func() {
		if err := app.Listen(":" + PORT); err != nil {
			log.Printf("Server error: %v\n", err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c
	log.Println("Shutting down server...")
	app.Shutdown()
}

func setMiddleware(app *fiber.App) {
	app.Use(logger.New())
}

func setDefaultRouter(app *fiber.App) {
	app.
		Get("/ping", func(c *fiber.Ctx) error {
			return c.SendString("success")
		}).
		Get("/livenss", func(c *fiber.Ctx) error {
			return c.SendString("liveness")
		}).
		Get("/readiness", func(c *fiber.Ctx) error {
			return c.SendString("readiness")
		})
}

func setRouters(app *fiber.App) {
	//
}
