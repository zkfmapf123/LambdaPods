package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zkfmapf123/lambda-pods/cmd/examples/dto"
	"github.com/zkfmapf123/lambda-pods/cmd/examples/services"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type SettingHandler struct {
	settingService *services.SettingService
	logger         *zap.Logger
}

func NewSettingHandler(db *gorm.DB, logger *zap.Logger) *SettingHandler {
	return &SettingHandler{
		settingService: services.NewSettingService(db, logger),
		logger:         logger,
	}
}

// UpdateAssumeRoleARN godoc
// @Summary Update Assume Role ARN
// @Description Update AWS Assume Role ARN for settings
// @Tags settings
// @Accept json
// @Produce json
// @Param request body dto.SettingRequest true "Setting Request"
// @Success 200 {string} string "success"
// @Failure 400 {string} string "Bad Request"
// @Router /settings [post]
func (h SettingHandler) UpdateAssumeRoleARN(c *fiber.Ctx) error {
	var req dto.SettingRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).SendString("Invalid request body")
	}

	h.logger.Info("UpdateAssumeRoleARN called", zap.String("assume_role_arn", req.AssumeRoleARN))

	return c.SendString("Assume Role ARN updated successfully")
}
