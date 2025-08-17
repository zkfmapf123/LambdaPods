package domains

type Settings struct {
	DBHost        string `gorm:"type:varchar(255);not null" json:"db_host"`
	ExternalID    string `gorm:"type:varchar(255);not null" json:"external_id"`
	AssumeRoleARN string `gorm:"type:varchar(255)" json:"assume_role_arn"`
}

func (s *Settings) TableName() string {
	return "settings"
}

func (s *Settings) IsEmptyDBHost() bool {
	return s.DBHost == ""
}

func (s *Settings) IsMatchDBHost(dbHost string) bool {
	return s.DBHost == dbHost
}
