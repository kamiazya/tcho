package link

import (
	"github.com/kamiazya/tcho/core/domain/model"
	"github.com/kamiazya/tcho/domain/repository"
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

func TitleLike(title string) SearchOption {
	return func(stmp *Stmp) error {
		if stmp.TitleLike != nil {
			return repository.ErrOptionAlreadySeted
		}
		stmp.TitleLike = &title
		return nil
	}
}
