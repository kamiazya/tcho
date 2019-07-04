package series

import (
	"time"

	"github.com/kamiazya/tcho/core/domain/model/tag"
	"github.com/kamiazya/tcho/core/domain/value/colorcode"
	"github.com/kamiazya/tcho/core/domain/value/memo"
	"github.com/kamiazya/tcho/core/domain/model"
	"github.com/kamiazya/tcho/core/domain/model/attachment"
	"github.com/kamiazya/tcho/core/domain/model/category"
	"github.com/kamiazya/tcho/core/domain/model/location"
	"github.com/kamiazya/tcho/core/domain/model/note"
	"github.com/kamiazya/tcho/core/domain/model/task"
)

type Series struct {
	model.ID

	// Setting
	Type Type

	WeekSeq  [7]bool
	MonthSeq [6]bool

	Since         time.Time
	Until         *time.Time
	ExceptionDate []time.Time

	// Schedule

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
