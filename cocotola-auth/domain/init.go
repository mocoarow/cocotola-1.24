package domain

import (
	"errors"
)

const AppName = "cocotola-auth"

var (
	ErrUnauthenticated = errors.New("unauthenticated")
)
