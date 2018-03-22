package note

import (
	"time"

	"bitbucket.org/kamiazya/tcho/core/domain/model"
)

type Stmp struct {
	ByIDs         []model.ID
	ByTitle       *string
	Text          *string
	HasLink       *bool
	LinkIDs       []model.ID
	HasTag        *bool
	Tags          []model.ID
	TagsLike      *string
	HasAttachment *bool
	AttachmentIDs []model.ID
	CreatedSince  *time.Time
	CreatedUntil  *time.Time
	Limit         *uint
	Offset        *uint
}
