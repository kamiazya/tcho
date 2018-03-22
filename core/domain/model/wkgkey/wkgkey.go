package wkgkey

import (
	"time"

	"bitbucket.org/kamiazya/tcho/core/domain/value/colorcode"
	"bitbucket.org/kamiazya/tcho/core/domain/value/memo"
	"bitbucket.org/kamiazya/tcho/core/domain/model"
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
