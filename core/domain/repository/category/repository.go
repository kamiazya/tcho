package category

import (
	"context"

	"github.com/kamiazya/tcho/core/domain/model"
	"github.com/kamiazya/tcho/core/domain/model/category"
)

type Repository interface {
	One(ctx context.Context, opts ...SearchOption) (*category.Category, error)
	List(ctx context.Context, opts ...SearchOption) ([]*category.Category, error)
	Stream(ctx context.Context, opts ...SearchOption) (<-chan *category.Category, error)

	Update(ctx context.Context, t *category.Category) (err error)
	Store(ctx context.Context, t *category.Category) (ID model.ID, err error)
	Delete(ctx context.Context, IDs ...model.ID) (success bool, err error)
}
