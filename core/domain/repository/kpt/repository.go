package kpt

import (
	"context"

	"github.com/kamiazya/tcho/core/domain/model"
	"github.com/kamiazya/tcho/core/domain/model/kpt"
)

type Repository interface {
	One(ctx context.Context, opts ...SearchOption) (*kpt.Kpt, error)
	List(ctx context.Context, opts ...SearchOption) ([]*kpt.Kpt, error)
	Stream(ctx context.Context, opts ...SearchOption) (<-chan *kpt.Kpt, error)

	Update(ctx context.Context, l *kpt.Kpt) (err error)
	Store(ctx context.Context, l *kpt.Kpt) (ID model.ID, err error)
	Delete(IDs ...model.ID) (success bool, err error)
}
