package location

import (
	"bitbucket.org/kamiazya/tcho/core/domain/value/memo"
	"bitbucket.org/kamiazya/tcho/core/domain/model"
)

type Location struct {
	model.ID
	Name string

	Description memo.Memo
}
