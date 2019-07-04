package note

import (
	"time"

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

func Text(text string) SearchOption {
	return func(stmp *Stmp) error {
		if stmp.Text != nil {
			return repository.ErrOptionAlreadySeted
		}
		stmp.Text = &text
		return nil
	}
}

func HasAttachment(hasAttachment bool) SearchOption {
	return func(stmp *Stmp) error {
		if stmp.HasAttachment != nil {
			return repository.ErrOptionAlreadySeted
		}
		stmp.HasAttachment = &hasAttachment
		return nil
	}
}

func AttachmentIDs(IDs []model.ID) SearchOption {
	return func(stmp *Stmp) error {
		if stmp.AttachmentIDs != nil {
			return repository.ErrOptionAlreadySeted
		}
		stmp.AttachmentIDs = IDs
		return nil
	}
}

func AttachmentID(ID model.ID) SearchOption {
	return func(stmp *Stmp) error {
		if stmp.AttachmentIDs != nil {
			return repository.ErrOptionAlreadySeted
		}
		stmp.AttachmentIDs = []model.ID{ID}
		return nil
	}
}

func HasLink(hasLink bool) SearchOption {
	return func(stmp *Stmp) error {
		if stmp.HasLink != nil {
			return repository.ErrOptionAlreadySeted
		}
		stmp.HasLink = &hasLink
		return nil
	}
}

func LinkIDs(IDs []model.ID) SearchOption {
	return func(stmp *Stmp) error {
		if stmp.LinkIDs != nil {
			return repository.ErrOptionAlreadySeted
		}
		stmp.LinkIDs = IDs
		return nil
	}
}

func LinkID(ID model.ID) SearchOption {
	return func(stmp *Stmp) error {
		if stmp.LinkIDs != nil {
			return repository.ErrOptionAlreadySeted
		}
		stmp.LinkIDs = []model.ID{ID}
		return nil
	}
}

func Tags(IDs []model.ID) SearchOption {
	return func(stmp *Stmp) error {
		if stmp.Tags != nil {
			return repository.ErrOptionAlreadySeted
		}
		stmp.Tags = IDs
		return nil
	}
}

func Tag(ID model.ID) SearchOption {
	return func(stmp *Stmp) error {
		if stmp.Tags != nil {
			return repository.ErrOptionAlreadySeted
		}
		stmp.Tags = []model.ID{ID}
		return nil
	}
}

func TagsLike(name string) SearchOption {
	return func(stmp *Stmp) error {
		if stmp.TagsLike != nil {
			return repository.ErrOptionAlreadySeted
		}
		stmp.TagsLike = &name
		return nil
	}
}

func CreatedSince(since time.Time) SearchOption {
	return func(stmp *Stmp) error {
		if stmp.CreatedSince != nil {
			return repository.ErrOptionAlreadySeted
		}
		stmp.CreatedSince = &since
		return nil
	}
}

func CreatedUntil(until time.Time) SearchOption {
	return func(stmp *Stmp) error {
		if stmp.CreatedUntil != nil {
			return repository.ErrOptionAlreadySeted
		}
		stmp.CreatedUntil = &until
		return nil
	}
}

func CreatedBetween(since, until time.Time) SearchOption {
	return func(stmp *Stmp) error {
		if stmp.CreatedSince != nil && stmp.CreatedUntil != nil {
			return repository.ErrOptionAlreadySeted
		}
		stmp.CreatedSince = &since
		stmp.CreatedUntil = &until
		return nil
	}
}

func Limit(limit uint) SearchOption {
	return func(stmp *Stmp) error {
		if stmp.Limit != nil {
			return repository.ErrOptionAlreadySeted
		}
		stmp.Limit = &limit
		return nil
	}
}

func Offset(offset uint) SearchOption {
	return func(stmp *Stmp) error {
		if stmp.Offset != nil {
			return repository.ErrOptionAlreadySeted
		}
		stmp.Offset = &offset
		return nil
	}
}
