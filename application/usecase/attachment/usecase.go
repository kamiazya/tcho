package attachment

import (
	"context"
	"io"

	"bitbucket.org/kamiazya/tcho/application/usecase"
	"bitbucket.org/kamiazya/tcho/domain/model"
	"bitbucket.org/kamiazya/tcho/domain/model/attachment"
	repository "bitbucket.org/kamiazya/tcho/domain/repository/attachment"
	"bitbucket.org/kamiazya/tcho/infrastructure/database"
)

// New returns UseCase.
func New(container *database.Container) (uc UseCase, err error) {
	if container.Attachment == nil {
		return nil, usecase.ErrRepositoryNotPrepared
	}
	uc = &useCase{
		repo: container.Attachment,
	}
	return
}

// UseCase is usecase of attachments.
type UseCase interface {
	GetAttachmentsAll(ctx context.Context) ([]*attachment.Attachment, error)
	SaveAttachment(ctx context.Context, a *attachment.Attachment, file io.Reader) (ID model.ID, err error)
}

var _ UseCase = new(useCase)

type useCase struct {
	repo repository.Repository
}

func (u *useCase) SaveAttachment(ctx context.Context, a *attachment.Attachment, file io.Reader) (ID model.ID, err error) {
	ID, err = u.repo.Store(ctx, a, file)
	return
}

func (u *useCase) GetAttachmentsAll(ctx context.Context) (attachments []*attachment.Attachment, err error) {
	attachments, err = u.repo.List(ctx)
	return
}

// func (u *useCase) RemoveAttachment(ctx context.Context, id model.ID) (sccuess bool, err error) {
// 	sccuess, err = u.repo.Delete(ctx, id)
// 	return
// }

// func (u *useCase) UpdateDescription(ctx context.Context, id model.ID, t value.TextType, text string) {

// }
