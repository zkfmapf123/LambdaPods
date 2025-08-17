package middlewares

import (
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/zkfmapf123/lambda-pods/cmd/examples/dto"
)

func ValidateMiddleware[T any]() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req T

		// bind json
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(dto.NewResponse[any](nil, fiber.StatusBadRequest, err.Error()))
		}

		// validate
		if err := validator.New().Struct(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(dto.NewResponse[any](nil, fiber.StatusBadRequest, err.Error()))
		}

		return c.Next()
	}
}
