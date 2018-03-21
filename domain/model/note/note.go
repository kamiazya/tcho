package note

import (
	"time"

	"bitbucket.org/kamiazya/tcho/domain/model"
	"bitbucket.org/kamiazya/tcho/domain/model/attachment"
	"bitbucket.org/kamiazya/tcho/domain/model/link"
	"bitbucket.org/kamiazya/tcho/domain/model/tag"
	"bitbucket.org/kamiazya/tcho/domain/value/memo"
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
