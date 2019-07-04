package schedule

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

func StartSince(since time.Time) SearchOption {
	return func(stmp *Stmp) error {
		if stmp.StartSince != nil {
			return repository.ErrOptionAlreadySeted
		}
		stmp.StartSince = &since
		return nil
	}
}

func StartUntil(start time.Time) SearchOption {
	return func(stmp *Stmp) error {
		if stmp.StartUntil != nil {
			return repository.ErrOptionAlreadySeted
		}
		stmp.StartUntil = &start
		return nil
	}
}

func StartBetween(from, until time.Time) SearchOption {
	return func(stmp *Stmp) error {
		if stmp.StartSince != nil && stmp.StartUntil != nil {
			return repository.ErrOptionAlreadySeted
		}
		stmp.StartSince = &from
		stmp.StartUntil = &until
		return nil
	}
}

func EndSince(since time.Time) SearchOption {
	return func(stmp *Stmp) error {
		if stmp.EndSince != nil {
			return repository.ErrOptionAlreadySeted
		}
		stmp.EndSince = &since
		return nil
	}
}

func EndUntil(end time.Time) SearchOption {
	return func(stmp *Stmp) error {
		if stmp.EndUntil != nil {
			return repository.ErrOptionAlreadySeted
		}
		stmp.EndUntil = &end
		return nil
	}
}

func EndBetween(from, until time.Time) SearchOption {
	return func(stmp *Stmp) error {
		if stmp.EndSince != nil && stmp.EndUntil != nil {
			return repository.ErrOptionAlreadySeted
		}
		stmp.EndSince = &from
		stmp.EndUntil = &until
		return nil
	}
}

func HasTask(hasTask bool) SearchOption {
	return func(stmp *Stmp) error {
		if stmp.HasTask != nil {
			return repository.ErrOptionAlreadySeted
		}
		stmp.HasTask = &hasTask
		return nil
	}
}

func TaskIDs(IDs []model.ID) SearchOption {
	return func(stmp *Stmp) error {
		if stmp.TaskIDs != nil {
			return repository.ErrOptionAlreadySeted
		}
		stmp.TaskIDs = IDs
		return nil
	}
}

func TaskID(ID model.ID) SearchOption {
	return func(stmp *Stmp) error {
		if stmp.TaskIDs != nil {
			return repository.ErrOptionAlreadySeted
		}
		stmp.TaskIDs = []model.ID{ID}
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

func NoteIDs(noteIDs []model.ID) SearchOption {
	return func(stmp *Stmp) error {
		if stmp.NoteIDs != nil {
			return repository.ErrOptionAlreadySeted
		}
		stmp.NoteIDs = noteIDs
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
