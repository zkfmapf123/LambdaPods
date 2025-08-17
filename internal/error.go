package internal

import "errors"

// Database
var (
	ErrInvalidUserRole       = errors.New("invalid user role, role must be readonly, developer, admin")
	ErrSettingEmptyDBHost    = errors.New("setting.DBHost is empty")
	ErrSettingNotMatchDBHost = errors.New("setting.DBHost is not match")
)

// Authentication
var ErrInvalidRole = errors.New("you are not authorized to access this resource")
