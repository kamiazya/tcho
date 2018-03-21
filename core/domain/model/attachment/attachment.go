package attachment

import (
	"time"

	"bitbucket.org/kamiazya/tcho/domain/model"
	"bitbucket.org/kamiazya/tcho/domain/model/tag"
	"bitbucket.org/kamiazya/tcho/domain/value/memo"
)

type Attachment struct {
	model.ID
	Name        string
	Mime        string
	Description memo.Memo
	Tags        []*tag.Tag

	Created time.Time
}
