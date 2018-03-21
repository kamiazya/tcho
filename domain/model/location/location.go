package location

import (
	"bitbucket.org/kamiazya/tcho/domain/model"
	"bitbucket.org/kamiazya/tcho/domain/value/memo"
)

type Location struct {
	model.ID
	Name string

	Description memo.Memo
}
