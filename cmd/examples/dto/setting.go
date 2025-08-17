package dto

type SettingAssumeRoleRequest struct {
	AssumeRoleARN string `json:"assume_role_arn" validate:"required" example:"arn:aws:iam::123456789012:role/MyRole"`
}

type SettingOpenAPIKeyRequest struct {
	OpenAPIKey string `json:"openapi_key" validate:"required" example:"sk-1234567890abcdef"`
}
