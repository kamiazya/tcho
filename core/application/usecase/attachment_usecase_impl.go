package usecase

import (
	"context"
	"io"

	"github.com/kamiazya/tcho/core/domain/model"
	"github.com/kamiazya/tcho/core/domain/model/attachment"
	"github.com/kamiazya/tcho/core/domain/repository"
)

var _ AttachmentUseCase = new(attachmentUseCase)

type attachmentUseCase struct {
	repo repository.AttachmentRepository
}

func (u *attachmentUseCase) SaveAttachment(ctx context.Context, a *attachment.Attachment, file io.Reader) (ID model.ID, err error) {
	ID, err = u.repo.Store(ctx, a, file)
	return
}

func (u *attachmentUseCase) GetAttachmentsAll(ctx context.Context) (attachments []*attachment.Attachment, err error) {
	attachments, err = u.repo.List(ctx)
	return
}
