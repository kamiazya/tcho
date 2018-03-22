package tag

import (
	"time"

	"bitbucket.org/kamiazya/tcho/core/domain/value/colorcode"
	"bitbucket.org/kamiazya/tcho/core/domain/value/memo"
	"bitbucket.org/kamiazya/tcho/core/domain/model"
)

type Tag struct {
	model.ID
	Name        string
	Color       colorcode.ColorCode
	Description memo.Memo
	Created     time.Time
}
