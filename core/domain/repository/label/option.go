package tag

import (
	"bitbucket.org/kamiazya/tcho/domain/model"
	"bitbucket.org/kamiazya/tcho/domain/repository"
)

type SearchOption func(*Stmp) error

func By(t model.ModelType, id model.ID) SearchOption {
	return func(stmp *Stmp) error {
		if stmp.Type != nil && stmp.ModelIDs != nil {
			return repository.ErrOptionAlreadySeted
		}
		stmp.Type = &t
		stmp.ModelIDs = []model.ID{id}
		return nil
	}
}

func Maintainer(maintainer string) SearchOption {
	return func(stmp *Stmp) error {
		if stmp.Maintainer != nil {
			return repository.ErrOptionAlreadySeted
		}
		stmp.Maintainer = &maintainer
		return nil
	}
}

func ValueLike(valueLike string) SearchOption {
	return func(stmp *Stmp) error {
		if stmp.ValueLike != nil {
			return repository.ErrOptionAlreadySeted
		}
		stmp.ValueLike = &valueLike
		return nil
	}
}
