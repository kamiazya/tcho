package pubsub

import "bitbucket.org/kamiazya/tcho/domain/model"

type Event struct {
	model.ModelType
	model.ID
	EventType
}
