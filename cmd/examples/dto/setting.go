package dto

type SettingRequest struct {
	AssumeRoleARN string `json:"assume_role_arn" validate:"required"`
}
