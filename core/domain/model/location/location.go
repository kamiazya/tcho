package location

import (
	"github.com/kamiazya/tcho/core/domain/value/memo"
	"github.com/kamiazya/tcho/core/domain/model"
)

type Location struct {
	model.ID
	Name string

	Description memo.Memo
}
