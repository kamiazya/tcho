package tag

import "bitbucket.org/kamiazya/tcho/domain/model"

type Stmp struct {
	Type       *model.ModelType
	ModelIDs   []model.ID
	Maintainer *string
	ValueLike  *string
}
