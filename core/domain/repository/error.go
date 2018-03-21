package repository

import "errors"

var (
	ErrOptionAlreadySeted = errors.New("ErrOptionAlreadySeted")

	ErrNotFound = errors.New("ErrNotFound")

	ErrUpdateTargetNotExist = errors.New("ErrUpdateTargetNotExist")

	ErrAlreadyHasID = errors.New("ErrAlreadyHasID")
)
