package pubsub

import "bitbucket.org/kamiazya/tcho/core/domain/model"

type Event struct {
	model.ModelType
	model.ID
	EventType
}
