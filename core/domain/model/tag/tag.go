package tag

import (
	"time"

	"github.com/kamiazya/tcho/core/domain/value/colorcode"
	"github.com/kamiazya/tcho/core/domain/value/memo"
	"github.com/kamiazya/tcho/core/domain/model"
)

type Tag struct {
	model.ID
	Name        string
	Color       colorcode.ColorCode
	Description memo.Memo
	Created     time.Time
}
