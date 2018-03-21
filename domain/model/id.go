package model

type ID interface {
	ID() string
	IsNew() bool
}
