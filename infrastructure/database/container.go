package database

import (
	"bitbucket.org/kamiazya/tcho/domain/repository/attachment"
	"bitbucket.org/kamiazya/tcho/domain/repository/category"
	"bitbucket.org/kamiazya/tcho/domain/repository/kpt"
	"bitbucket.org/kamiazya/tcho/domain/repository/link"
	"bitbucket.org/kamiazya/tcho/domain/repository/location"
	"bitbucket.org/kamiazya/tcho/domain/repository/note"
	"bitbucket.org/kamiazya/tcho/domain/repository/schedule"
	"bitbucket.org/kamiazya/tcho/domain/repository/tag"
	"bitbucket.org/kamiazya/tcho/domain/repository/task"
	"bitbucket.org/kamiazya/tcho/domain/repository/wkgkey"
	"bitbucket.org/kamiazya/tcho/domain/repository/wkgvalue"
)

type Container struct {
	Attachment   attachment.Repository
	Category     category.Repository
	Kpt          kpt.Repository
	Link         link.Repository
	Location     location.Repository
	Note         note.Repository
	Schedule     schedule.Repository
	Tag          tag.Repository
	Task         task.Repository
	WorkingKey   wkgkey.Repository
	WorkingValue wkgvalue.Repository
}
