package pubsub

import "github.com/kamiazya/tcho/core/domain/model"

type Event struct {
	model.ModelType
	model.ID
	EventType
}
