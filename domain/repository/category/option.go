package category

import (
	"bitbucket.org/kamiazya/tcho/domain/model"
	"bitbucket.org/kamiazya/tcho/domain/model/category"
	"bitbucket.org/kamiazya/tcho/domain/repository"
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

func NameLike(name string) SearchOption {
	return func(stmp *Stmp) error {
		if stmp.NameLike != nil {
			return repository.ErrOptionAlreadySeted
		}
		stmp.NameLike = &name
		return nil
	}
}

func ByType(t category.Type) SearchOption {
	return func(stmp *Stmp) error {
		if stmp.ByType != nil {
			return repository.ErrOptionAlreadySeted
		}
		stmp.ByType = &t
		return nil
	}
}
