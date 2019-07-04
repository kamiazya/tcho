package note

import (
	"time"

	"github.com/kamiazya/tcho/core/domain/model/tag"
	"github.com/kamiazya/tcho/core/domain/value/memo"
	"github.com/kamiazya/tcho/core/domain/model"
	"github.com/kamiazya/tcho/core/domain/model/attachment"
	"github.com/kamiazya/tcho/core/domain/model/link"
)

type Note struct {
	model.ID
	Title string
	Body  memo.Memo

	Links       []link.Link
	Tags        []tag.Tag
	Attachments []attachment.Attachment

	Created  time.Time
	Modified time.Time
}
