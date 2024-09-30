package service

import "errors"

var (
	ErrGenerateTokens    = errors.New("failed to generate tokens")
	ErrUserNotFound      = errors.New("user was not found")
	ErrPasswordsNotMatch = errors.New("passwords do not match")
	ErrHashingPassword   = errors.New("failed to hash password")
)
