package task

import (
	"time"

	"bitbucket.org/kamiazya/tcho/domain/model"
	"bitbucket.org/kamiazya/tcho/domain/model/category"
	"bitbucket.org/kamiazya/tcho/domain/model/location"
	"bitbucket.org/kamiazya/tcho/domain/model/tag"
	"bitbucket.org/kamiazya/tcho/domain/model/wkgvalue"
	"bitbucket.org/kamiazya/tcho/domain/value/colorcode"
	"bitbucket.org/kamiazya/tcho/domain/value/memo"
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
