package location

import (
	"context"

	"bitbucket.org/kamiazya/tcho/core/domain/model"
	"bitbucket.org/kamiazya/tcho/core/domain/model/location"
)

type Repository interface {
	One(ctx context.Context, opts ...SearchOption) (*location.Location, error)
	List(ctx context.Context, opts ...SearchOption) ([]*location.Location, error)
	Stream(ctx context.Context, opts ...SearchOption) (<-chan *location.Location, error)

	Update(ctx context.Context, l *location.Location) (err error)
	Store(ctx context.Context, l *location.Location) (ID model.ID, err error)
	Delete(IDs ...model.ID) (success bool, err error)
}
