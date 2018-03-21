package pubsub

import (
	"context"
)

type Server interface {
	Emmit(ctx context.Context, event Event)
	Serve(ctx context.Context)
	MarshalEvent(e Event) []byte
}
