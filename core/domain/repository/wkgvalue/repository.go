package wkgvalue

import (
	"context"

	"bitbucket.org/kamiazya/tcho/core/domain/model"
	"bitbucket.org/kamiazya/tcho/core/domain/model/wkgvalue"
)

type Repository interface {
	One(ctx context.Context, opts ...SearchOption) (*wkgvalue.WorkingValue, error)
	List(ctx context.Context, opts ...SearchOption) ([]*wkgvalue.WorkingValue, error)
	Stream(ctx context.Context, opts ...SearchOption) (<-chan *wkgvalue.WorkingValue, error)

	Update(ctx context.Context, l *wkgvalue.WorkingValue) (err error)
	Store(ctx context.Context, l *wkgvalue.WorkingValue) (ID model.ID, err error)
	Delete(ctx context.Context, IDs ...model.ID) (success bool, err error)
}
