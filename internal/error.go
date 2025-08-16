package internal

import "errors"

// Database
var ErrInvalidUserRole = errors.New("invalid user role, role must be readonly, developer, admin")
