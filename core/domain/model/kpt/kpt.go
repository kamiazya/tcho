package kpt

import (
	"time"

	"github.com/kamiazya/tcho/core/domain/value/memo"
	"github.com/kamiazya/tcho/core/domain/model"
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
