package sqlite

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"
	"io"
	"io/ioutil"
	"strings"
	"time"

	"bitbucket.org/kamiazya/tcho/domain/model"
	"bitbucket.org/kamiazya/tcho/domain/model/attachment"
	"bitbucket.org/kamiazya/tcho/domain/model/tag"
	repository "bitbucket.org/kamiazya/tcho/domain/repository/attachment"
	"bitbucket.org/kamiazya/tcho/domain/value/memo"
	"bitbucket.org/kamiazya/tcho/infrastructure/database"
)

var _ repository.Repository = new(attachmentRepository)

type attachmentRepository struct {
	db      *sql.DB
	tagRepo *tagRepository
}

type attachmentSchema struct {
	ID              intID          `sql:"id"`
	Name            string         `sql:"name"`
	Mime            sql.NullString `sql:"mime"`
	DescriptionType sql.NullInt64  `sql:"description_type"`
	DescriptionText sql.NullString `sql:"description_text"`
	Created         time.Time      `sql:"created"`
}

func (repo *attachmentRepository) encode(a *attachment.Attachment) (row *attachmentSchema) {
	row = new(attachmentSchema)

	// ID
	if id, ok := a.ID.(*intID); ok {
		row.ID = *id
	}

	// // Name
	row.Name = a.Name

	if a.Description.Type != memo.None {
		row.DescriptionType.Scan(a.Description.Type)
		row.DescriptionText.Scan(a.Description.Text)
	}

	// Mime
	if a.Mime != "" {
		row.Mime.Scan(a.Mime)
	}

	// Created
	row.Created = a.Created

	return
}

func (repo *attachmentRepository) decode(row *attachmentSchema) (a *attachment.Attachment) {

	a = new(attachment.Attachment)

	a.ID = &row.ID
	a.Name = row.Name

	// Mime
	if row.Mime.Valid {
		a.Mime = row.Mime.String
	}

	// Description
	if row.DescriptionType.Valid && row.DescriptionText.Valid {
		a.Description.Type = memo.TextType(row.DescriptionType.Int64)
		a.Description.Text = row.DescriptionText.String
	}

	return
}

func (repo *attachmentRepository) selectQuery(opts ...repository.SearchOption) (query string, args []interface{}, err error) {
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
      mime,
      description_type,
      description_text,
      created
    FROM
      tbl_attachment
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

func (repo *attachmentRepository) One(ctx context.Context, opts ...repository.SearchOption) (*attachment.Attachment, error) {
	opts = append(opts, repository.Limit(1))
	query, args, queryErr := repo.selectQuery(opts...)
	if queryErr != nil {
		return nil, queryErr
	}

	var err error
	a := new(attachment.Attachment)
	row := repo.db.QueryRowContext(ctx, query, args...)
	err = row.Scan(
		&a.ID,
		&a.Name,
		&a.Mime,
		&a.Description.Type,
		&a.Description.Text,
		&a.Created,
	)
	if err != nil {
		return nil, err
	}

	if a.Tags, err = repo.tags(ctx, a.ID); err != nil {
		return nil, err
	}
	return a, nil
}

func (repo *attachmentRepository) tags(ctx context.Context, attachmentID model.ID) (tags []*tag.Tag, err error) {

	rows, err := repo.db.QueryContext(
		ctx,
		`SELECT
      id,
      name,
      color,
      description_type,
      description_text,
      created
    FROM
      tbl_tag AS tag
    INNER JOIN tbl_attachment_tag_relation AS relation
      ON
        relation.attachment_id = ?
          AND
        relation.tag_id = tag.id
    `,
		attachmentID.ID(),
	)
	if err != nil {
		return nil, err
	}

	tags = make([]*tag.Tag, 0)

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
		tags = append(tags, repo.tagRepo.decode(t))
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tags, nil
}

func (repo *attachmentRepository) File(ctx context.Context, ID model.ID) (reader io.Reader, err error) {
	var bs []byte

	query :=
		`SELECT
      blob
    FROM
      tbl_attachment
    WHERE id = ?
  `
	row := repo.db.QueryRowContext(ctx, query, ID.ID())
	if err = row.Scan(&bs); err != nil {
		return nil, err
	}
	return bytes.NewReader(bs), nil
}

func (repo *attachmentRepository) List(ctx context.Context, opts ...repository.SearchOption) ([]*attachment.Attachment, error) {
	query, args, queryErr := repo.selectQuery(opts...)
	if queryErr != nil {
		return nil, queryErr
	}
	rows, err := repo.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	attachments := make([]*attachment.Attachment, 0)
	for rows.Next() {
		a := new(attachmentSchema)
		scanErr := rows.Scan(
			&a.ID,
			&a.Name,
			&a.Mime,
			&a.DescriptionType,
			&a.DescriptionText,
			&a.Created,
		)
		if scanErr != nil {
			return nil, scanErr
		}
		attachments = append(attachments, repo.decode(a))
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return attachments, nil
}

func (repo *attachmentRepository) Update(ctx context.Context, t *attachment.Attachment) error {
	res, err := repo.db.ExecContext(ctx,
		`UPADTE tbl_attachment
      SET
        name = ?,
        mime = ?,
        description_type = ?,
        description_text = ?,
        created = ?
      WHERE
        id = ?`,
		t.Name,
		t.Mime,
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

func (repo *attachmentRepository) Store(ctx context.Context, a *attachment.Attachment, r io.Reader) (ID model.ID, err error) {
	row := repo.encode(a)

	blob, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	res, err := repo.db.ExecContext(ctx,
		`INSERT INTO tbl_attachment (
      name,
      mime,
      description_type,
      description_text,
      created,
      blob
    )VALUES ( ?, ?, ?, ?, ?, ?)`,
		row.Name,
		row.Mime,
		row.DescriptionType,
		row.DescriptionText,
		row.Created,
		blob,
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

func (repo *attachmentRepository) Delete(ctx context.Context, IDs ...model.ID) (success bool, err error) {
	var query string
	var arg string
	switch len(IDs) {
	case 0:
		err = database.ErrIdNotSet
	case 1:
		query = "DELETE FROM tbl_attachment WHERE id = ?"
		arg = IDs[0].ID()
	default:
		query = "DELETE FROM tbl_attachment WHERE id IN (?)"
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
