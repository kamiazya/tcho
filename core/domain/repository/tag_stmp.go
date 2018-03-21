package repository

import "bitbucket.org/kamiazya/tcho/core/domain/model"

type Stmp struct {
	ByIDs    []model.ID
	NameLike *string
	Limit    *uint
	Offset   *uint
}
