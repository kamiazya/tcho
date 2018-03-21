package category

import (
	"time"

	"bitbucket.org/kamiazya/tcho/domain/model"
	"bitbucket.org/kamiazya/tcho/domain/model/wkgkey"
	"bitbucket.org/kamiazya/tcho/domain/value/colorcode"
	"bitbucket.org/kamiazya/tcho/domain/value/memo"
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
