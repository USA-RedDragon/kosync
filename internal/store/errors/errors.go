package errors

import "errors"

var (
	ErrUserNotFound     = errors.New("user not found")
	ErrProgressNotFound = errors.New("progress not found")
)
