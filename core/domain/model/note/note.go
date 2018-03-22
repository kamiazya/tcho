package note

import (
	"time"

	"bitbucket.org/kamiazya/tcho/core/domain/model/tag"
	"bitbucket.org/kamiazya/tcho/core/domain/value/memo"
	"bitbucket.org/kamiazya/tcho/core/domain/model"
	"bitbucket.org/kamiazya/tcho/core/domain/model/attachment"
	"bitbucket.org/kamiazya/tcho/core/domain/model/link"
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
