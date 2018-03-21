package tag

import (
	"time"

	"bitbucket.org/kamiazya/tcho/domain/model"
	"bitbucket.org/kamiazya/tcho/domain/value/colorcode"
	"bitbucket.org/kamiazya/tcho/domain/value/memo"
)

type Tag struct {
	model.ID
	Name        string
	Color       colorcode.ColorCode
	Description memo.Memo
	Created     time.Time
}
