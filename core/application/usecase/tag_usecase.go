package usecase

import (
	"context"

	"bitbucket.org/kamiazya/tcho/core/domain/model"
	"bitbucket.org/kamiazya/tcho/core/domain/model/tag"
	"bitbucket.org/kamiazya/tcho/core/infrastructure/database"
)

// NewTagUseCase returns NewTagUseCase.
func NewTagUseCase(container *database.Container) (uc TagUseCase, err error) {
	if container.Tag == nil {
		return nil, ErrRepositoryNotPrepared
	}
	uc = &tagUseCase{
		repo: container.Tag,
	}
	return
}

// TagUseCase is usecase of tags.
type TagUseCase interface {
	GetTagsAll(ctx context.Context) ([]*tag.Tag, error)
	AddTag(ctx context.Context, t *tag.Tag) (ID model.ID, err error)
}
