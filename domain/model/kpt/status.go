package kpt

type Status uint

const (
	Unconfirmed Status = iota
	NotExecuted
	Executed
)
