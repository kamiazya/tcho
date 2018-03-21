package kpt

import (
	"bitbucket.org/kamiazya/tcho/domain/model"
	"bitbucket.org/kamiazya/tcho/domain/model/kpt"
)

type Stmp struct {
	ByIDs  []model.ID
	Type   *kpt.Type
	Status *kpt.Status
}
