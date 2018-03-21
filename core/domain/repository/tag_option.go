package repository

import "bitbucket.org/kamiazya/tcho/core/domain/model"

type TagSearchOption func(*TagStmp) error

func ById(id model.ID) TagSearchOption {
	return func(stmp *TagStmp) error {
		if stmp.ByIDs != nil {
			return ErrOptionAlreadySeted
		}
		stmp.ByIDs = []model.ID{id}
		return nil
	}
}

func ByIds(ids ...model.ID) TagSearchOption {
	return func(stmp *TagStmp) error {
		if stmp.ByIDs != nil {
			return ErrOptionAlreadySeted
		}
		stmp.ByIDs = ids
		return nil
	}
}

func NameLike(name string) TagSearchOption {
	return func(stmp *TagStmp) error {
		if stmp.NameLike != nil {
			return ErrOptionAlreadySeted
		}
		stmp.NameLike = &name
		return nil
	}
}

func Limit(limit uint) TagSearchOption {
	return func(stmp *TagStmp) error {
		if stmp.Limit != nil {
			return ErrOptionAlreadySeted
		}
		stmp.Limit = &limit
		return nil
	}
}

func Offset(offset uint) TagSearchOption {
	return func(stmp *TagStmp) error {
		if stmp.Offset != nil {
			return ErrOptionAlreadySeted
		}
		stmp.Offset = &offset
		return nil
	}
}
