package category

import (
	"time"

	"bitbucket.org/kamiazya/tcho/core/domain/value/colorcode"
	"bitbucket.org/kamiazya/tcho/core/domain/value/memo"
	"bitbucket.org/kamiazya/tcho/core/domain/model"
	"bitbucket.org/kamiazya/tcho/core/domain/model/wkgkey"
)

type Category struct {
	model.ID
	Type Type

	Name  string
	Color colorcode.ColorCode

	Description memo.Memo

	DefaultWorkingKeys []wkgkey.WorkingKey

	Created  time.Time
	Modified time.Time
}
