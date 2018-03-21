package kpt

import (
	"time"

	"bitbucket.org/kamiazya/tcho/domain/model"
	"bitbucket.org/kamiazya/tcho/domain/value/memo"
)

type Kpt struct {
	model.ID
	Type   Type
	Status Status

	// Prarent
	Content memo.Memo

	Created  time.Time
	Modified time.Time
}
