package tag

import "github.com/kamiazya/tcho/core/domain/model"

type Stmp struct {
	Type       *model.ModelType
	ModelIDs   []model.ID
	Maintainer *string
	ValueLike  *string
}
