package kpt

import (
	"github.com/kamiazya/tcho/core/domain/model"
	"github.com/kamiazya/tcho/core/domain/model/kpt"
)

type Stmp struct {
	ByIDs  []model.ID
	Type   *kpt.Type
	Status *kpt.Status
}
