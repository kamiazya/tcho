package task

import (
	"time"

	"bitbucket.org/kamiazya/tcho/core/domain/model/tag"
	"bitbucket.org/kamiazya/tcho/core/domain/value/colorcode"
	"bitbucket.org/kamiazya/tcho/core/domain/value/memo"
	"bitbucket.org/kamiazya/tcho/core/domain/model"
	"bitbucket.org/kamiazya/tcho/core/domain/model/category"
	"bitbucket.org/kamiazya/tcho/core/domain/model/location"
	"bitbucket.org/kamiazya/tcho/core/domain/model/wkgvalue"
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
