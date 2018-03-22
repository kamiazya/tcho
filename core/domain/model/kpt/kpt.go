package kpt

import (
	"time"

	"bitbucket.org/kamiazya/tcho/core/domain/value/memo"
	"bitbucket.org/kamiazya/tcho/core/domain/model"
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
