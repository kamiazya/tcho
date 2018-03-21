package database

import "errors"

var (
	ErrNoRowsAffected = errors.New("ErrNoRowsAffected")
	ErrIdNotSet       = errors.New("ErrIdNotSet")
)
