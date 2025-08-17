package services

import (
	"github.com/google/uuid"
	"github.com/zkfmapf123/lambda-pods/cmd/examples/repository"
	"github.com/zkfmapf123/lambda-pods/internal"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type SettingService struct {
	settingRepository *repository.SettingsRepository
	logger            *zap.Logger
}

func NewSettingService(db *gorm.DB, logger *zap.Logger) *SettingService {
	return &SettingService{
		settingRepository: repository.NewSettingRepository(db),
		logger:            logger,
	}
}

func (s *SettingService) InitSetting(dbHost, setRandID string) {
	err := s.settingRepository.SelectHost(dbHost)

	newExternalID := uuid.New().String()

	// DB 없을 경우
	if err == internal.ErrSettingEmptyDBHost {
		s.logger.Info("setting.DBHost is empty", zap.String("dbHost", dbHost), zap.String("externalID", newExternalID))
		s.settingRepository.Create(dbHost, newExternalID)
		return
	}

	// DB Host가 맞지 않을 경우
	if err == internal.ErrSettingNotMatchDBHost {
		s.logger.Info("setting.DBHost is changed", zap.String("dbHost", dbHost), zap.String("externalID", newExternalID))
		s.settingRepository.Update(dbHost, newExternalID)
		return
	}

	s.logger.Info("setting.DBHost is match", zap.String("dbHost", dbHost), zap.String("externalID", newExternalID))
}
