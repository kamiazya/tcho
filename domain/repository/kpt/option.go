package kpt

import (
	"bitbucket.org/kamiazya/tcho/domain/model"
	"bitbucket.org/kamiazya/tcho/domain/model/kpt"
	"bitbucket.org/kamiazya/tcho/domain/repository"
)

var (
	// Status
	Unconfirmed = Status(kpt.Unconfirmed)
	NotExecuted = Status(kpt.NotExecuted)
	Executed    = Status(kpt.Executed)

	// Type
	Keep    = Type(kpt.Keep)
	Problem = Type(kpt.Problem)
	Try     = Type(kpt.Try)
)

type SearchOption func(*Stmp) error

func ById(id model.ID) SearchOption {
	return func(stmp *Stmp) error {
		if stmp.ByIDs != nil {
			return repository.ErrOptionAlreadySeted
		}
		stmp.ByIDs = []model.ID{id}
		return nil
	}
}

func ByIds(ids ...model.ID) SearchOption {
	return func(stmp *Stmp) error {
		if stmp.ByIDs != nil {
			return repository.ErrOptionAlreadySeted
		}
		stmp.ByIDs = ids
		return nil
	}
}

func Type(t kpt.Type) SearchOption {
	return func(stmp *Stmp) error {
		if stmp.Type != nil {
			return repository.ErrOptionAlreadySeted
		}
		stmp.Type = &t
		return nil
	}
}

func Status(s kpt.Status) SearchOption {
	return func(stmp *Stmp) error {
		if stmp.Status != nil {
			return repository.ErrOptionAlreadySeted
		}
		stmp.Status = &s
		return nil
	}
}
