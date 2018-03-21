package pubsub

type EventType uint

const (
	_ EventType = iota

	Created
	Updated
	Deleted
)
