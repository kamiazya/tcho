package model

type ModelType uint

const (
	Attachment ModelType = iota
	Category
	Kpt
	Label
	Link
	Location
	Note
	Project
	Schedule
	Tag
	Task
	WorkingKey
	WorkingValue
)

type ModelTypes []ModelType

func (ms ModelTypes) Exist(ts ...ModelType) (exit bool) {
	for _, m := range ms {
		for _, t := range ts {
			if m == t {
				return true
			}
		}
	}
	return exit
}
