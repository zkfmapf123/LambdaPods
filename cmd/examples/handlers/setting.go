package handlers

import (
	"github.com/gofiber/fiber/v2"
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

func (h SettingHandler) UpdateAssumeRoleARN(c *fiber.Ctx) error {
	return c.SendString("test")
}
