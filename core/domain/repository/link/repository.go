package link

import (
	"context"

	"github.com/kamiazya/tcho/core/domain/model"
	"github.com/kamiazya/tcho/core/domain/model/link"
)

type Repository interface {
	One(ctx context.Context, opts ...SearchOption) (*link.Link, error)
	List(ctx context.Context, opts ...SearchOption) ([]*link.Link, error)
	Stream(ctx context.Context, opts ...SearchOption) (<-chan *link.Link, error)

	Update(ctx context.Context, l *link.Link) (err error)
	Store(ctx context.Context, l *link.Link) (ID model.ID, err error)
	Delete(ctx context.Context, IDs ...model.ID) (success bool, err error)
}
