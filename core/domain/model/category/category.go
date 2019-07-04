package category

import (
	"time"

	"github.com/kamiazya/tcho/core/domain/value/colorcode"
	"github.com/kamiazya/tcho/core/domain/value/memo"
	"github.com/kamiazya/tcho/core/domain/model"
	"github.com/kamiazya/tcho/core/domain/model/wkgkey"
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
