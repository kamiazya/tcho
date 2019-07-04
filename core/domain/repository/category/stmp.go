package category

import (
	"github.com/kamiazya/tcho/core/domain/model"
	"github.com/kamiazya/tcho/core/domain/model/category"
)

type Stmp struct {
	ByIDs    []model.ID
	ByType   *category.Type
	NameLike *string
}
