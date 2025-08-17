package repository

import (
	"github.com/zkfmapf123/lambda-pods/domains"
	"github.com/zkfmapf123/lambda-pods/internal"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type SettingsRepository struct {
	db     *gorm.DB
	logger *zap.Logger
}

func NewSettingRepository(db *gorm.DB) *SettingsRepository {
	return &SettingsRepository{
		db: db,
	}
}

func (s SettingsRepository) SelectHost(dbHost string) error {
	setting := domains.Settings{}

	s.db.Where("db_host = ?", dbHost).Find(&setting)

	if setting.IsEmptyDBHost() {
		return internal.ErrSettingEmptyDBHost
	}

	if !setting.IsMatchDBHost(dbHost) {
		return internal.ErrSettingNotMatchDBHost
	}

	return nil
}

func (s SettingsRepository) Create(dbHost string, externalID string) error {
	setting := domains.Settings{
		DBHost:     dbHost,
		ExternalID: externalID,
	}

	s.db.Create(&setting)
	return nil
}

func (s SettingsRepository) Update(dbHost, externalID string) error {
	setting := domains.Settings{
		DBHost:     dbHost,
		ExternalID: externalID,
	}

	s.db.Save(&setting)
	return nil
}
