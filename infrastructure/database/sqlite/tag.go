package sqlite

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"bitbucket.org/kamiazya/tcho/domain/model"
	"bitbucket.org/kamiazya/tcho/domain/model/tag"
	repository "bitbucket.org/kamiazya/tcho/domain/repository/tag"
	"bitbucket.org/kamiazya/tcho/domain/value/colorcode"
	"bitbucket.org/kamiazya/tcho/domain/value/memo"
	"bitbucket.org/kamiazya/tcho/infrastructure/database"
)

var _ repository.Repository = new(tagRepository)

type tagRepository struct {
	db *sql.DB
}

type tagSchema struct {
	ID              intID          `sql:"id"`
	Name            string         `sql:"name"`
	Color           sql.NullInt64  `sql:"color"`
	DescriptionType sql.NullInt64  `sql:"description_type"`
	DescriptionText sql.NullString `sql:"description_text"`
	Created         time.Time      `sql:"created"`
}

func (repo *tagRepository) encode(t *tag.Tag) (row *tagSchema) {
	row = new(tagSchema)

	// ID
	if id, ok := t.ID.(*intID); ok {
		row.ID = *id
	}

	// Name
	row.Name = t.Name

	if t.Description.Type != memo.None {
		row.DescriptionType.Scan(t.Description.Type)
		row.DescriptionText.Scan(t.Description.Text)
	}

	// Color
	if t.Color != colorcode.NoColor {
		row.Color.Scan(t.Color)
	}

	// Created
	row.Created = t.Created

	return
}

func (repo *tagRepository) decode(row *tagSchema) (t *tag.Tag) {

	t = new(tag.Tag)

	t.ID = &row.ID
	t.Name = row.Name

	if row.Color.Valid {
		t.Color = colorcode.ColorCode(row.Color.Int64)
	}

	// Description
	if row.DescriptionType.Valid && row.DescriptionText.Valid {
		t.Description.Type = memo.TextType(row.DescriptionType.Int64)
		t.Description.Text = row.DescriptionText.String
	}

	return
}

func (repo *tagRepository) selectQuery(opts ...repository.SearchOption) (query string, args []interface{}, err error) {
	stmp := new(repository.Stmp)
	for _, opt := range opts {
		if err = opt(stmp); err != nil {
			return "", nil, err
		}
	}

	conditions := make([]string, 0, len(opts))
	args = make([]interface{}, 0, len(opts))
	if stmp.ByIDs != nil && len(stmp.ByIDs) != 0 {
		if len(stmp.ByIDs) == 1 {
			conditions = append(conditions, "`id` = ?")
			args = append(args, stmp.ByIDs[0])
		} else {
			conditions = append(conditions, "`id` IN (?)")
			args = append(args, stmp.ByIDs)
		}
	}

	if stmp.NameLike != nil && *stmp.NameLike != "" {
		conditions = append(conditions, "`name` IN (?)")
		args = append(args, *stmp.NameLike)
	}

	queryBuf := bytes.NewBufferString(`
    SELECT
      id,
      name,
      color,
      description_type,
      description_text,
      created
    FROM
      tbl_tag
  `)
	if len(conditions) > 0 {
		queryBuf.WriteString(" WHERE ")
		queryBuf.WriteString(strings.Join(conditions, " AND "))
	}

	if stmp.Limit != nil && *stmp.Limit > 0 {
		queryBuf.WriteString(" LIMIT ")
		queryBuf.WriteString(fmt.Sprint(stmp.Limit))
	}
	if stmp.Offset != nil && *stmp.Offset > 0 {
		queryBuf.WriteString(" OFFSET ")
		queryBuf.WriteString(fmt.Sprint(stmp.Offset))
	}

	return queryBuf.String(), args, nil
}

func (repo *tagRepository) One(ctx context.Context, opts ...repository.SearchOption) (*tag.Tag, error) {
	opts = append(opts, repository.Limit(1))
	query, args, queryErr := repo.selectQuery(opts...)
	if queryErr != nil {
		return nil, queryErr
	}

	t := new(tag.Tag)
	row := repo.db.QueryRowContext(ctx, query, args...)
	scanErr := row.Scan(
		&t.ID,
		&t.Name,
		&t.Color,
		&t.Description.Type,
		&t.Description.Text,
		&t.Created,
	)
	if scanErr != nil {
		return nil, scanErr
	}
	return t, nil
}

func (repo *tagRepository) List(ctx context.Context, opts ...repository.SearchOption) ([]*tag.Tag, error) {
	query, args, queryErr := repo.selectQuery(opts...)
	if queryErr != nil {
		return nil, queryErr
	}
	rows, err := repo.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	tags := make([]*tag.Tag, 0)
	for rows.Next() {
		t := new(tagSchema)
		scanErr := rows.Scan(
			&t.ID,
			&t.Name,
			&t.Color,
			&t.DescriptionType,
			&t.DescriptionText,
			&t.Created,
		)
		if scanErr != nil {
			return nil, scanErr
		}
		tags = append(tags, repo.decode(t))
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tags, nil
}

func (repo *tagRepository) Update(ctx context.Context, t *tag.Tag) error {
	res, err := repo.db.ExecContext(ctx,
		`UPADTE tbl_tag
      SET
        name = ?,
        color = ?,
        description_type = ?,
        description_text = ?,
        created = ?
      WHERE
        id = ?`,
		t.Name,
		t.Color,
		t.Description.Type,
		t.Description.Text,
		t.Created,
		t.ID.ID(),
	)
	if err != nil {
		return err
	}

	if affected, _ := res.RowsAffected(); affected == 0 {
		return database.ErrNoRowsAffected
	}
	return nil
}

func (repo *tagRepository) Store(ctx context.Context, t *tag.Tag) (ID model.ID, err error) {
	row := repo.encode(t)

	res, err := repo.db.ExecContext(ctx,
		`INSERT INTO tbl_tag (
      name,
      color,
      description_type,
      description_text,
      created
    )VALUES ( ?, ?, ?, ?, ?)`,
		row.Name,
		row.Color,
		row.DescriptionType,
		row.DescriptionText,
		row.Created,
	)
	if err != nil {
		return nil, err
	}

	lastInsertID, execErr := res.LastInsertId()
	if execErr != nil {
		return nil, execErr
	}

	id := intID(lastInsertID)
	return &id, nil
}

func (repo *tagRepository) Delete(ctx context.Context, IDs ...model.ID) (success bool, err error) {
	var query string
	var arg string
	switch len(IDs) {
	case 0:
		err = database.ErrIdNotSet
	case 1:
		query = "DELETE FROM tbl_tag WHERE id = ?"
		arg = IDs[0].ID()
	default:
		query = "DELETE FROM tbl_tag WHERE id IN (?)"
		ids := make([]string, 0, len(IDs))
		for _, id := range IDs {
			ids = append(ids, id.ID())
		}
		arg = strings.Join(ids, ", ")
	}

	res, err := repo.db.ExecContext(ctx, query, arg)

	if affected, _ := res.RowsAffected(); affected == 0 {
		err = database.ErrNoRowsAffected
	} else {
		success = true
	}
	return
}
