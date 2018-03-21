package link

import (
	"net/url"
	"time"

	"bitbucket.org/kamiazya/tcho/domain/model"
	"bitbucket.org/kamiazya/tcho/domain/model/tag"
	"bitbucket.org/kamiazya/tcho/domain/value/memo"
)

type Link struct {
	model.ID
	Title       string
	Description memo.Memo
	URL         url.URL
	Tags        []tag.Tag
	Created     time.Time
}
