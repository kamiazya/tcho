package attachment

import (
	"context"
	"io"

	"bitbucket.org/kamiazya/tcho/domain/model"
	"bitbucket.org/kamiazya/tcho/domain/model/attachment"
)

type Repository interface {
	One(ctx context.Context, opts ...SearchOption) (*attachment.Attachment, error)
	List(ctx context.Context, opts ...SearchOption) ([]*attachment.Attachment, error)

	File(ctx context.Context, ID model.ID) (reader io.Reader, err error)

	Update(ctx context.Context, a *attachment.Attachment) (err error)
	Store(ctx context.Context, a *attachment.Attachment, r io.Reader) (ID model.ID, err error)
	Delete(ctx context.Context, IDs ...model.ID) (success bool, err error)
}
