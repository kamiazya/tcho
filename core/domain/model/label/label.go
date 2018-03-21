package label

import (
	"time"

	"bitbucket.org/kamiazya/tcho/domain/model"
)

type Label struct {
	model.ID

	ModelType  model.ModelType
	ModelID    model.ID
	Value      string
	Maintainer string
	Created    time.Time
}
