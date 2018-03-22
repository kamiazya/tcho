package note

import (
	"context"

	"bitbucket.org/kamiazya/tcho/core/domain/model"
	"bitbucket.org/kamiazya/tcho/core/domain/model/schedule"
)

type Repository interface {
	One(ctx context.Context, opts ...SearchOption) (*schedule.Schedule, error)
	List(ctx context.Context, opts ...SearchOption) ([]*schedule.Schedule, error)
	Stream(ctx context.Context, opts ...SearchOption) (<-chan *schedule.Schedule, error)

	Update(ctx context.Context, s *schedule.Schedule) (err error)
	Store(ctx context.Context, s *schedule.Schedule) (ID model.ID, err error)
	Delete(IDs ...model.ID) (success bool, err error)
}
