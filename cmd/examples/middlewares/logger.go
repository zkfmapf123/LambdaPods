package middlewares

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/zkfmapf123/lambda-pods/internal"
	"go.uber.org/zap"
)

func LoggingMiddleware(l *zap.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()

		if internal.GetProcessEnv()["ENV"] != "dev" {
			// request
			l.Info("HTTP Request Started",
				zap.String("method", c.Method()),
				zap.String("path", c.Path()),
				zap.String("ip", c.IP()),
				zap.String("user-agent", c.Get("User-Agent")),
			)
		}

		err := c.Next()

		// response
		if internal.GetProcessEnv()["ENV"] != "dev" {
			l.Info("HTTP Response Completed",
				zap.String("method", c.Method()),
				zap.String("path", c.Path()),
				zap.String("ip", c.IP()),
				zap.Int("status", c.Response().StatusCode()),
				zap.Duration("latency", time.Since(start)),
			)
		}

		return err
	}
}
