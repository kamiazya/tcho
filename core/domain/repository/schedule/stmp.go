package schedule

import (
	"time"

	"github.com/kamiazya/tcho/core/domain/model"
)

type Stmp struct {
	ByIDs         []model.ID
	ByTitle       *string
	StartSince    *time.Time
	StartUntil    *time.Time
	EndSince      *time.Time
	EndUntil      *time.Time
	HasTask       *bool
	TaskIDs       []model.ID
	HasNote       *bool
	NoteIDs       []model.ID
	HasAttachment *bool
	AttachmentIDs []model.ID
	Tags          []model.ID
	TagsLike      *string
	Limit         *uint
	Offset        *uint
}
