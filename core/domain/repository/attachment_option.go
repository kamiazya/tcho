package repository

import (
	"bitbucket.org/kamiazya/tcho/core/domain/model"
)

type AttachmentSearchOption func(*AttachmentStmp) error

func AttachmentById(id model.ID) AttachmentSearchOption {
	return func(stmp *AttachmentStmp) error {
		if stmp.ByIDs != nil {
			return ErrOptionAlreadySeted
		}
		stmp.ByIDs = []model.ID{id}
		return nil
	}
}

func AttachmentByIds(ids ...model.ID) AttachmentSearchOption {
	return func(stmp *AttachmentStmp) error {
		if stmp.ByIDs != nil {
			return ErrOptionAlreadySeted
		}
		stmp.ByIDs = ids
		return nil
	}
}

func AttachmentNameLike(name string) AttachmentSearchOption {
	return func(stmp *AttachmentStmp) error {
		if stmp.NameLike != nil {
			return ErrOptionAlreadySeted
		}
		stmp.NameLike = &name
		return nil
	}
}

func AttachmentLimit(limit uint) AttachmentSearchOption {
	return func(stmp *AttachmentStmp) error {
		if stmp.Limit != nil {
			return ErrOptionAlreadySeted
		}
		stmp.Limit = &limit
		return nil
	}
}

func AttachmentOffset(offset uint) AttachmentSearchOption {
	return func(stmp *AttachmentStmp) error {
		if stmp.Offset != nil {
			return ErrOptionAlreadySeted
		}
		stmp.Offset = &offset
		return nil
	}
}
