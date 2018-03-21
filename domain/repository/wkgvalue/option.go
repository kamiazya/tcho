package wkgvalue

import (
	"bitbucket.org/kamiazya/tcho/domain/model"
	"bitbucket.org/kamiazya/tcho/domain/repository"
)

type SearchOption func(*Stmp) error

func Key(id model.ID) SearchOption {
	return func(stmp *Stmp) error {
		if stmp.Key != nil {
			return repository.ErrOptionAlreadySeted
		}
		stmp.Key = id
		return nil
	}
}

func Task(id model.ID) SearchOption {
	return func(stmp *Stmp) error {
		if stmp.Task != nil {
			return repository.ErrOptionAlreadySeted
		}
		stmp.Task = id
		return nil
	}
}
