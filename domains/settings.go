package domains

type Settings struct {
	ProjectID     string `gorm:"type:varchar(255);not null;unique" json:"project_id"`
	ExternalID    string `gorm:"type:varchar(255);not null;unique" json:"external_id"`
	AssumeRoleARN string `gorm:"type:varchar(255);not null" json:"assume_role_arn"`
}
