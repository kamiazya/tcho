package label

import (
	"time"

	"github.com/kamiazya/tcho/core/domain/model"
)

type Label struct {
	model.ID

	ModelType  model.ModelType
	ModelID    model.ID
	Value      string
	Maintainer string
	Created    time.Time
}
