package repository

import (
	"context"
	"io"

	"github.com/kamiazya/tcho/core/domain/model"
	"github.com/kamiazya/tcho/core/domain/model/attachment"
)

type AttachmentRepository interface {
	One(ctx context.Context, opts ...AttachmentSearchOption) (*attachment.Attachment, error)
	List(ctx context.Context, opts ...AttachmentSearchOption) ([]*attachment.Attachment, error)

	File(ctx context.Context, ID model.ID) (reader io.Reader, err error)

	Update(ctx context.Context, a *attachment.Attachment) (err error)
	Store(ctx context.Context, a *attachment.Attachment, r io.Reader) (ID model.ID, err error)
	Delete(ctx context.Context, IDs ...model.ID) (success bool, err error)
}
