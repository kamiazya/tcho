package attachment

import (
	"time"

	"github.com/kamiazya/tcho/core/domain/model/tag"
	"github.com/kamiazya/tcho/core/domain/value/memo"
	"github.com/kamiazya/tcho/core/domain/model"
)

type Attachment struct {
	model.ID
	Name        string
	Mime        string
	Description memo.Memo
	Tags        []*tag.Tag

	Created time.Time
}
