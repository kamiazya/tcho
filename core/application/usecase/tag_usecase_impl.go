package usecase

import (
	"context"

	"github.com/kamiazya/tcho/core/domain/model"
	"github.com/kamiazya/tcho/core/domain/model/tag"
	"github.com/kamiazya/tcho/core/domain/repository"
	"github.com/kamiazya/tcho/core/domain/value/memo"
)

var _ TagUseCase = new(tagUseCase)

type tagUseCase struct {
	repo repository.TagRepository
}

func (u *tagUseCase) AddTag(ctx context.Context, t *tag.Tag) (ID model.ID, err error) {
	ID, err = u.repo.Store(ctx, t)
	return
}

func (u *tagUseCase) GetTagsAll(ctx context.Context) (tags []*tag.Tag, err error) {
	tags, err = u.repo.List(ctx)
	return
}

func (u *tagUseCase) RemoveTag(ctx context.Context, id model.ID) (sccuess bool, err error) {
	sccuess, err = u.repo.Delete(ctx, id)
	return
}

func (u *tagUseCase) UpdateDescription(ctx context.Context, id model.ID, t memo.TextType, text string) {

}
