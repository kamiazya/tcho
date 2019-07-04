package sqlite

import (
	"database/sql"

	"github.com/kamiazya/tcho/core/domain/model"
	"github.com/kamiazya/tcho/core/infrastructure/database"
)

func RepositoryContainer(db *sql.DB, requireds ...model.ModelType) (container *database.Container) {
	container = new(database.Container)

	var (
		tagRepo        *tagRepository
		attachmentRepo *attachmentRepository
		ms             = model.ModelTypes(requireds)
	)

	if ms.Exist(model.Tag, model.Attachment) {
		tagRepo = &tagRepository{
			db: db,
		}
		container.Tag = tagRepo
	}

	if ms.Exist(model.Attachment) {
		attachmentRepo = &attachmentRepository{
			db:      db,
			tagRepo: tagRepo,
		}
		container.Attachment = attachmentRepo
	}

	return
}
