package dto

// SettingAssumeRoleRequest represents the request to update AWS Assume Role ARN
type SettingAssumeRoleRequest struct {
	AssumeRoleARN string `json:"assume_role_arn" validate:"required" example:"arn:aws:iam::123456789012:role/MyRole"`
}

// SettingOpenAPIKeyRequest represents the request to update OpenAPI Key
type SettingOpenAPIKeyRequest struct {
	OpenAPIKey string `json:"openapi_key" validate:"required" example:"sk-1234567890abcdef"`
}
