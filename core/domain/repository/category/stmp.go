package category

import (
	"bitbucket.org/kamiazya/tcho/core/domain/model"
	"bitbucket.org/kamiazya/tcho/core/domain/model/category"
)

type Stmp struct {
	ByIDs    []model.ID
	ByType   *category.Type
	NameLike *string
}
