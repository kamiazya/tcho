package tag

import "bitbucket.org/kamiazya/tcho/domain/model"

type Stmp struct {
	ByIDs    []model.ID
	NameLike *string
	Limit    *uint
	Offset   *uint
}
