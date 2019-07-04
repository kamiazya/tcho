package pubsub

import (
	"context"
	"log"
	"time"

	"github.com/kamiazya/tcho/application/event/pubsub"
	"github.com/go-mangos/mangos"
)

func New(opts ...ServerOption) (pubsub.Server, error) {
	c := new(Config)
	for _, opt := range opts {
		if err := opt(c); err != nil {
			return nil, err
		}
	}

	return genFromConfig(c)
}

type nanomsgServer struct {
	url  string
	sock mangos.Socket
	ch   chan pubsub.Event
}

func (n *nanomsgServer) Emmit(ctx context.Context, event pubsub.Event) {
	n.ch <- event
}

func (n *nanomsgServer) Serve(ctx context.Context) {

	go n.mainLoop(ctx)
}

func (n *nanomsgServer) mainLoop(ctx context.Context) {
	var err error
	for {
		select {
		case e := <-n.ch:
			// Could also use sock.RecvMsg to get header
			if err = n.sock.Send(n.MarshalEvent(e)); err != nil {
				log.Printf("Failed publishing: %s", err.Error())
			}
		case <-ctx.Done():
			break
		}
	}
}

func (n *nanomsgServer) MarshalEvent(e pubsub.Event) []byte {
	d := time.Now().Format(time.ANSIC)
	return []byte(d)
}
