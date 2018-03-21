package repository

import (
	"bitbucket.org/kamiazya/tcho/core/domain/model"
	"bitbucket.org/kamiazya/tcho/core/domain/repository"
)

type AttachmentSearchOption func(*AttachmentStmp) error

func ById(id model.ID) AttachmentSearchOption {
	return func(stmp *AttachmentStmp) error {
		if stmp.ByIDs != nil {
			return repository.ErrOptionAlreadySeted
		}
		stmp.ByIDs = []model.ID{id}
		return nil
	}
}

func ByIds(ids ...model.ID) AttachmentSearchOption {
	return func(stmp *AttachmentStmp) error {
		if stmp.ByIDs != nil {
			return repository.ErrOptionAlreadySeted
		}
		stmp.ByIDs = ids
		return nil
	}
}

func NameLike(name string) AttachmentSearchOption {
	return func(stmp *AttachmentStmp) error {
		if stmp.NameLike != nil {
			return repository.ErrOptionAlreadySeted
		}
		stmp.NameLike = &name
		return nil
	}
}

func Limit(limit uint) AttachmentSearchOption {
	return func(stmp *AttachmentStmp) error {
		if stmp.Limit != nil {
			return repository.ErrOptionAlreadySeted
		}
		stmp.Limit = &limit
		return nil
	}
}

func Offset(offset uint) AttachmentSearchOption {
	return func(stmp *AttachmentStmp) error {
		if stmp.Offset != nil {
			return repository.ErrOptionAlreadySeted
		}
		stmp.Offset = &offset
		return nil
	}
}
