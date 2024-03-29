package task

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

func ByTitle(title string) SearchOption {
	return func(stmp *Stmp) error {
		if stmp.ByTitle != nil {
			return repository.ErrOptionAlreadySeted
		}
		stmp.ByTitle = &title
		return nil
	}
}

func InLocation(ID model.ID) SearchOption {
	return func(stmp *Stmp) error {
		if stmp.InLocation != nil {
			return repository.ErrOptionAlreadySeted
		}
		stmp.InLocation = ID
		return nil
	}
}

func DeadlineSince(since time.Time) SearchOption {
	return func(stmp *Stmp) error {
		if stmp.DeadlineSince != nil {
			return repository.ErrOptionAlreadySeted
		}
		stmp.DeadlineSince = &since
		return nil
	}
}

func DeadlineUntil(Deadline time.Time) SearchOption {
	return func(stmp *Stmp) error {
		if stmp.DeadlineUntil != nil {
			return repository.ErrOptionAlreadySeted
		}
		stmp.DeadlineUntil = &Deadline
		return nil
	}
}

func DeadlineBetween(from, until time.Time) SearchOption {
	return func(stmp *Stmp) error {
		if stmp.DeadlineSince != nil && stmp.DeadlineUntil != nil {
			return repository.ErrOptionAlreadySeted
		}
		stmp.DeadlineSince = &from
		stmp.DeadlineUntil = &until
		return nil
	}
}

func NoteIDs(IDs []model.ID) SearchOption {
	return func(stmp *Stmp) error {
		if stmp.NoteIDs != nil {
			return repository.ErrOptionAlreadySeted
		}
		stmp.NoteIDs = IDs
		return nil
	}
}

func NoteID(ID model.ID) SearchOption {
	return func(stmp *Stmp) error {
		if stmp.NoteIDs != nil {
			return repository.ErrOptionAlreadySeted
		}
		stmp.NoteIDs = []model.ID{ID}
		return nil
	}
}

func HasNote(hasNote bool) SearchOption {
	return func(stmp *Stmp) error {
		if stmp.HasNote != nil {
			return repository.ErrOptionAlreadySeted
		}
		stmp.HasNote = &hasNote
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
