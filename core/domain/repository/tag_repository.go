package repository

import (
	"context"

	"bitbucket.org/kamiazya/tcho/core/domain/model"
	"bitbucket.org/kamiazya/tcho/core/domain/model/tag"
)

type TagRepository interface {
	One(ctx context.Context, opts ...TagSearchOption) (*tag.Tag, error)
	List(ctx context.Context, opts ...TagSearchOption) ([]*tag.Tag, error)

	Update(ctx context.Context, t *tag.Tag) (err error)
	Store(ctx context.Context, t *tag.Tag) (ID model.ID, err error)
	Delete(ctx context.Context, IDs ...model.ID) (success bool, err error)
}
