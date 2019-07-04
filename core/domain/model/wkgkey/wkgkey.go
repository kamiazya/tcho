package wkgkey

import (
	"time"

	"github.com/kamiazya/tcho/core/domain/value/colorcode"
	"github.com/kamiazya/tcho/core/domain/value/memo"
	"github.com/kamiazya/tcho/core/domain/model"
)

type WorkingKey struct {
	model.ID
	Type        Type
	Name        string
	Description memo.Memo
	Color       colorcode.ColorCode
	Created     time.Time
	Modified    time.Time
}
