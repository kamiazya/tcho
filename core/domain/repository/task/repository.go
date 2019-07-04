package task

import (
	"context"

	"github.com/kamiazya/tcho/core/domain/model"
	"github.com/kamiazya/tcho/core/domain/model/task"
)

type Repository interface {
	One(ctx context.Context, opts ...SearchOption) (*task.Task, error)
	List(ctx context.Context, opts ...SearchOption) ([]*task.Task, error)
	Stream(ctx context.Context, opts ...SearchOption) (<-chan *task.Task, error)

	Update(ctx context.Context, s *task.Task) (err error)
	Store(ctx context.Context, s *task.Task) (ID model.ID, err error)
	Delete(IDs ...model.ID) (success bool, err error)
}
