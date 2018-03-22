package tag

import (
	"context"

	"bitbucket.org/kamiazya/tcho/core/domain/model"
	"bitbucket.org/kamiazya/tcho/core/domain/model/label"
)

type Repository interface {
	One(ctx context.Context, opts ...SearchOption) (*label.Label, error)
	List(ctx context.Context, opts ...SearchOption) ([]*label.Label, error)

	Update(ctx context.Context, t *label.Label) (err error)
	Store(ctx context.Context, t *label.Label) (ID model.ID, err error)
	Delete(ctx context.Context, IDs ...model.ID) (success bool, err error)
}
