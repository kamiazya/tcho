package link

import (
	"net/url"
	"time"

	"github.com/kamiazya/tcho/core/domain/model/tag"
	"github.com/kamiazya/tcho/core/domain/value/memo"
	"github.com/kamiazya/tcho/core/domain/model"
)

type Link struct {
	model.ID
	Title       string
	Description memo.Memo
	URL         url.URL
	Tags        []tag.Tag
	Created     time.Time
}
