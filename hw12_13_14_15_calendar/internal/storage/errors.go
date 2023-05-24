package storage

import "errors"

var (
	ErrEventNotFound = errors.New("event not found")
	ErrNoConnection  = errors.New("connection is not exist")
)
