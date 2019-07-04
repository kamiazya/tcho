package database

import "github.com/kamiazya/tcho/core/domain/repository"

type Container struct {
	Attachment repository.AttachmentRepository
	Tag        repository.TagRepository
}
