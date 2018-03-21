package category

import (
	"bitbucket.org/kamiazya/tcho/domain/model"
	"bitbucket.org/kamiazya/tcho/domain/model/category"
)

type Stmp struct {
	ByIDs    []model.ID
	ByType   *category.Type
	NameLike *string
}
