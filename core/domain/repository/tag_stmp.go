package repository

import "github.com/kamiazya/tcho/core/domain/model"

type TagStmp struct {
	ByIDs    []model.ID
	NameLike *string
	Limit    *uint
	Offset   *uint
}
