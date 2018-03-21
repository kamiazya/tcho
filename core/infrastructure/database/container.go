package database

import "bitbucket.org/kamiazya/tcho/core/domain/repository"

type Container struct {
	Attachment repository.AttachmentRepository
	Tag        repository.TagRepository
}
