package task

import (
	"time"

	"bitbucket.org/kamiazya/tcho/domain/model"
	"bitbucket.org/kamiazya/tcho/domain/model/task"
)

type Stmp struct {
	ByIDs         []model.ID
	ByTitle       *string
	ByType        *task.Type
	ByStatus      *task.Status
	InLocation    model.ID
	DeadlineSince *time.Time
	DeadlineUntil *time.Time
	Tags          []model.ID
	TagsLike      *string
	HasNote       *bool
	NoteIDs       []model.ID
	HasAttachment *bool
	AttachmentIDs []model.ID
	Limit         *uint
	Offset        *uint
}
