package sqlite

import (
	"strconv"

	"bitbucket.org/kamiazya/tcho/domain/model"
)

var _ model.ID = new(intID)

type intID int64

func (i *intID) ID() string {
	return strconv.Itoa(int(*i))
}

func (i *intID) IsNew() bool {
	return i == nil
}
