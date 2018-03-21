package tag

import (
	"context"

	"bitbucket.org/kamiazya/tcho/application/usecase"
	"bitbucket.org/kamiazya/tcho/domain/model"
	"bitbucket.org/kamiazya/tcho/domain/model/tag"
	repository "bitbucket.org/kamiazya/tcho/domain/repository/tag"
	"bitbucket.org/kamiazya/tcho/domain/value/memo"
	"bitbucket.org/kamiazya/tcho/infrastructure/database"
)

// New returns UseCase.
func New(container *database.Container) (uc UseCase, err error) {
	if container.Tag == nil {
		return nil, usecase.ErrRepositoryNotPrepared
	}
	uc = &useCase{
		repo: container.Tag,
	}
	return
}

// UseCase is usecase of tags.
type UseCase interface {
	GetTagsAll(ctx context.Context) ([]*tag.Tag, error)
	AddTag(ctx context.Context, t *tag.Tag) (ID model.ID, err error)
}

var _ UseCase = new(useCase)

type useCase struct {
	repo repository.Repository
}

func (u *useCase) AddTag(ctx context.Context, t *tag.Tag) (ID model.ID, err error) {
	ID, err = u.repo.Store(ctx, t)
	return
}

func (u *useCase) GetTagsAll(ctx context.Context) (tags []*tag.Tag, err error) {
	tags, err = u.repo.List(ctx)
	return
}

func (u *useCase) RemoveTag(ctx context.Context, id model.ID) (sccuess bool, err error) {
	sccuess, err = u.repo.Delete(ctx, id)
	return
}

func (u *useCase) UpdateDescription(ctx context.Context, id model.ID, t memo.TextType, text string) {

}
