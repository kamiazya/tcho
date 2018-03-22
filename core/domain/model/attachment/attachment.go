package attachment

import (
	"time"

	"bitbucket.org/kamiazya/tcho/core/domain/model/tag"
	"bitbucket.org/kamiazya/tcho/core/domain/value/memo"
	"bitbucket.org/kamiazya/tcho/core/domain/model"
)

type Attachment struct {
	model.ID
	Name        string
	Mime        string
	Description memo.Memo
	Tags        []*tag.Tag

	Created time.Time
}
