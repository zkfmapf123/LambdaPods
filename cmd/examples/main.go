// Package main Lambda Pods API
//
// This is a sample server for Lambda Pods.
//
// Terms Of Service: http://swagger.io/terms/
//
// Schemes: http, https
// Host: localhost:3000
// BasePath: /
// Version: 1.0.0
// License: MIT http://opensource.org/licenses/MIT
// Contact: your-email@example.com
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
//
// swagger:meta
package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger"
	"github.com/zkfmapf123/lambda-pods/cmd/examples/dto"
	"github.com/zkfmapf123/lambda-pods/cmd/examples/handlers"
	"github.com/zkfmapf123/lambda-pods/cmd/examples/middlewares"
	"github.com/zkfmapf123/lambda-pods/cmd/examples/services"
	"github.com/zkfmapf123/lambda-pods/configs"
	"github.com/zkfmapf123/lambda-pods/internal"
	"go.uber.org/zap"
	"gorm.io/gorm"

	_ "github.com/zkfmapf123/lambda-pods/cmd/examples/docs"
)

var (
	PORT           = internal.GetProcessEnv()["PORT"]
	APP_NAME       = internal.GetProcessEnv()["APP_NAME"]
	DB_HOST        = internal.GetProcessEnv()["DB_HOST"]
	DB_PORT        = internal.GetProcessEnv()["DB_PORT"]
	DB_USER        = internal.GetProcessEnv()["DB_USER"]
	DB_PASS        = internal.GetProcessEnv()["DB_PASS"]
	DB_NAME        = internal.GetProcessEnv()["DB_NAME"]
	SERVER_RAND_ID = internal.GetProcessEnv()["SERVER_RAND_ID"]
	ENV            = internal.GetProcessEnv()["ENV"]
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
	setMiddleware(app, logger)
	setRouters(app, RouterParams{
		db:     db,
		logger: logger,
	})

	go func() {
		if err := app.Listen(":" + PORT); err != nil {
			logger.Error("Server error", zap.Error(err))
		}
	}()

	/////////////////////////////////////////////// Init /////////////////////////////////////////////////
	InitDBHostAndExternalID(db, logger)
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

func setMiddleware(app *fiber.App, logger *zap.Logger) {
	app.Use(middlewares.LoggingMiddleware(logger))
}

func setDefaultRouter(app *fiber.App) {
	app.
		Get("/swagger/*", fiberSwagger.WrapHandler).
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
	settingHandler := handlers.NewSettingHandler(params.db, params.logger)

	// assumeRoleARN 설정
	app.Post("/settings", middlewares.ValidateMiddleware[dto.SettingRequest](), settingHandler.UpdateAssumeRoleARN)

	// 로그인
	// 유저 생성
	// 유저 삭제
	// 유저 Role 변경

	// Lambda 리스트 조회
	// Lambda 리스트 생성
	// Lambda 리스트 배포 (수정)
	// Lambda 리스트 삭제

	app.Get("/test", func(c *fiber.Ctx) error {
		params.logger.Info("test", zap.String("username", "leedonggyu"), zap.Int("age", 94), zap.String("job", "devops"))
		return c.SendString("test")
	})
}

// 시작할때 DB Host 변경
func InitDBHostAndExternalID(db *gorm.DB, logger *zap.Logger) {
	settingServce := services.NewSettingService(db, logger)
	settingServce.InitSetting(DB_HOST, SERVER_RAND_ID)
}
