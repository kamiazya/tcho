package wkgkey

import (
	"time"

	"bitbucket.org/kamiazya/tcho/domain/model"
	"bitbucket.org/kamiazya/tcho/domain/value/colorcode"
	"bitbucket.org/kamiazya/tcho/domain/value/memo"
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
