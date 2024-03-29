package repository

import "github.com/kamiazya/tcho/core/domain/model"

type TagSearchOption func(*TagStmp) error

func TagById(id model.ID) TagSearchOption {
	return func(stmp *TagStmp) error {
		if stmp.ByIDs != nil {
			return ErrOptionAlreadySeted
		}
		stmp.ByIDs = []model.ID{id}
		return nil
	}
}

func TagByIds(ids ...model.ID) TagSearchOption {
	return func(stmp *TagStmp) error {
		if stmp.ByIDs != nil {
			return ErrOptionAlreadySeted
		}
		stmp.ByIDs = ids
		return nil
	}
}

func TagNameLike(name string) TagSearchOption {
	return func(stmp *TagStmp) error {
		if stmp.NameLike != nil {
			return ErrOptionAlreadySeted
		}
		stmp.NameLike = &name
		return nil
	}
}

func TagLimit(limit uint) TagSearchOption {
	return func(stmp *TagStmp) error {
		if stmp.Limit != nil {
			return ErrOptionAlreadySeted
		}
		stmp.Limit = &limit
		return nil
	}
}

func TagOffset(offset uint) TagSearchOption {
	return func(stmp *TagStmp) error {
		if stmp.Offset != nil {
			return ErrOptionAlreadySeted
		}
		stmp.Offset = &offset
		return nil
	}
}
