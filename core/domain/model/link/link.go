package link

import (
	"net/url"
	"time"

	"bitbucket.org/kamiazya/tcho/core/domain/model/tag"
	"bitbucket.org/kamiazya/tcho/core/domain/value/memo"
	"bitbucket.org/kamiazya/tcho/core/domain/model"
)

type Link struct {
	model.ID
	Title       string
	Description memo.Memo
	URL         url.URL
	Tags        []tag.Tag
	Created     time.Time
}
