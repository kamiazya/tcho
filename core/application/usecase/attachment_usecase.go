package usecase

import (
	"context"
	"io"

	"github.com/kamiazya/tcho/core/domain/model"
	"github.com/kamiazya/tcho/core/domain/model/attachment"
	"github.com/kamiazya/tcho/core/infrastructure/database"
)

// NewAttachmentUseCase returns AttachmentUseCase.
func NewAttachmentUseCase(container *database.Container) (uc AttachmentUseCase, err error) {
	if container.Attachment == nil {
		return nil, ErrRepositoryNotPrepared
	}
	uc = &attachmentUseCase{
		repo: container.Attachment,
	}
	return
}

// AttachmentUseCase is usecase of attachments.
type AttachmentUseCase interface {
	GetAttachmentsAll(ctx context.Context) ([]*attachment.Attachment, error)
	SaveAttachment(ctx context.Context, a *attachment.Attachment, file io.Reader) (ID model.ID, err error)
}
