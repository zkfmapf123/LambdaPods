package internal

import "errors"

// Database
var ErrInvalidUserRole = errors.New("invalid user role, role must be readonly, developer, admin")

// Authentication
var ErrInvalidRole = errors.New("you are not authorized to access this resource")
