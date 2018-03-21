package schedule

import (
	"time"

	"bitbucket.org/kamiazya/tcho/domain/model"
	"bitbucket.org/kamiazya/tcho/domain/model/attachment"
	"bitbucket.org/kamiazya/tcho/domain/model/category"
	"bitbucket.org/kamiazya/tcho/domain/model/location"
	"bitbucket.org/kamiazya/tcho/domain/model/note"
	"bitbucket.org/kamiazya/tcho/domain/model/tag"
	"bitbucket.org/kamiazya/tcho/domain/model/task"
	"bitbucket.org/kamiazya/tcho/domain/value/colorcode"
	"bitbucket.org/kamiazya/tcho/domain/value/memo"
)

type Schedule struct {
	model.ID
	SeriesID    model.ID
	Description memo.Memo

	AllDay bool
	Start  time.Time
	End    time.Time

	Task task.Task

	Notes       []note.Note
	Attachments []attachment.Attachment
	CustomColor colorcode.ColorCode

	Category category.Category
	Tags     []tag.Tag

	Location location.Location

	Created  time.Time
	Modified time.Time
}

func (s Schedule) Duration() time.Duration {
	return s.End.Sub(s.Start)
}

func (s Schedule) IsSeries() bool {
	return s.SeriesID != nil
}
