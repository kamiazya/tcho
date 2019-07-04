package task

import (
	"time"

	"github.com/kamiazya/tcho/core/domain/model/tag"
	"github.com/kamiazya/tcho/core/domain/value/colorcode"
	"github.com/kamiazya/tcho/core/domain/value/memo"
	"github.com/kamiazya/tcho/core/domain/model"
	"github.com/kamiazya/tcho/core/domain/model/category"
	"github.com/kamiazya/tcho/core/domain/model/location"
	"github.com/kamiazya/tcho/core/domain/model/wkgvalue"
)

type Task struct {
	model.ID
	Title       string
	Type        Type
	Status      Status
	Description memo.Memo
	Deadline    time.Time
	SubTasks    []Task

	CustomColor   colorcode.ColorCode
	Category      category.Category
	WorkingValues []wkgvalue.WorkingValue
	Tags          []tag.Tag

	Location location.Location
	Created  time.Time
	Modified time.Time
}
