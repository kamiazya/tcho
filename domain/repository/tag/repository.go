package tag

import (
	"context"

	"bitbucket.org/kamiazya/tcho/domain/model"
	"bitbucket.org/kamiazya/tcho/domain/model/tag"
)

type Repository interface {
	One(ctx context.Context, opts ...SearchOption) (*tag.Tag, error)
	List(ctx context.Context, opts ...SearchOption) ([]*tag.Tag, error)

	Update(ctx context.Context, t *tag.Tag) (err error)
	Store(ctx context.Context, t *tag.Tag) (ID model.ID, err error)
	Delete(ctx context.Context, IDs ...model.ID) (success bool, err error)
}
