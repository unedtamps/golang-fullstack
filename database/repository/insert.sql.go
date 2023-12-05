// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: insert.sql

package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createAttachment = `-- name: CreateAttachment :one

INSERT INTO
    "attachment" (id, connection_id, id_messages)
VALUES ($1, $2, $3) RETURNING id,
    attachment_url
`

type CreateAttachmentParams struct {
	ID           string `db:"id"`
	ConnectionID string `db:"connection_id"`
	IDMessages   string `db:"id_messages"`
}

type CreateAttachmentRow struct {
	ID            string `db:"id"`
	AttachmentUrl string `db:"attachment_url"`
}

func (q *Queries) CreateAttachment(ctx context.Context, arg CreateAttachmentParams) (*CreateAttachmentRow, error) {
	row := q.db.QueryRow(ctx, createAttachment, arg.ID, arg.ConnectionID, arg.IDMessages)
	var i CreateAttachmentRow
	err := row.Scan(&i.ID, &i.AttachmentUrl)
	return &i, err
}

const createGroup = `-- name: CreateGroup :one

INSERT INTO
    "groups" (id, name, description)
VALUES ($1, $2, $3) RETURNING id
`

type CreateGroupParams struct {
	ID          string      `db:"id"`
	Name        string      `db:"name"`
	Description pgtype.Text `db:"description"`
}

func (q *Queries) CreateGroup(ctx context.Context, arg CreateGroupParams) (string, error) {
	row := q.db.QueryRow(ctx, createGroup, arg.ID, arg.Name, arg.Description)
	var id string
	err := row.Scan(&id)
	return id, err
}

const createGroupAttachement = `-- name: CreateGroupAttachement :one

INSERT INTO
    "group_attachment"(
        id,
        attachment_url,
        group_id,
        id_message
    )
VALUES ($1, $2, $3, $4) RETURNING id,
    attachment_url
`

type CreateGroupAttachementParams struct {
	ID            string `db:"id"`
	AttachmentUrl string `db:"attachment_url"`
	GroupID       string `db:"group_id"`
	IDMessage     string `db:"id_message"`
}

type CreateGroupAttachementRow struct {
	ID            string `db:"id"`
	AttachmentUrl string `db:"attachment_url"`
}

func (q *Queries) CreateGroupAttachement(ctx context.Context, arg CreateGroupAttachementParams) (*CreateGroupAttachementRow, error) {
	row := q.db.QueryRow(ctx, createGroupAttachement,
		arg.ID,
		arg.AttachmentUrl,
		arg.GroupID,
		arg.IDMessage,
	)
	var i CreateGroupAttachementRow
	err := row.Scan(&i.ID, &i.AttachmentUrl)
	return &i, err
}

const createGroupMessages = `-- name: CreateGroupMessages :one

INSERT INTO
    "group_messages" (
        id,
        user_id,
        group_id,
        message_content
    )
VALUES ($1, $2, $3, $4) RETURNING id
`

type CreateGroupMessagesParams struct {
	ID             pgtype.Text `db:"id"`
	UserID         string      `db:"user_id"`
	GroupID        string      `db:"group_id"`
	MessageContent string      `db:"message_content"`
}

func (q *Queries) CreateGroupMessages(ctx context.Context, arg CreateGroupMessagesParams) (pgtype.Text, error) {
	row := q.db.QueryRow(ctx, createGroupMessages,
		arg.ID,
		arg.UserID,
		arg.GroupID,
		arg.MessageContent,
	)
	var id pgtype.Text
	err := row.Scan(&id)
	return id, err
}

const createGroupParticiPants = `-- name: CreateGroupParticiPants :execrows

INSERT INTO "group_participants" VALUES($1,$2)
`

type CreateGroupParticiPantsParams struct {
	Column1 pgtype.Text `db:"column_1"`
	Column2 pgtype.Text `db:"column_2"`
}

func (q *Queries) CreateGroupParticiPants(ctx context.Context, arg CreateGroupParticiPantsParams) (int64, error) {
	result, err := q.db.Exec(ctx, createGroupParticiPants, arg.Column1, arg.Column2)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected(), nil
}

const createMessage = `-- name: CreateMessage :one

INSERT INTO
    "messages" (
        id,
        connection_id,
        message_content,
        user_id
    )
VALUES ($1, $2, $3, $4) RETURNING id
`

type CreateMessageParams struct {
	ID             pgtype.Text `db:"id"`
	ConnectionID   string      `db:"connection_id"`
	MessageContent pgtype.Text `db:"message_content"`
	UserID         string      `db:"user_id"`
}

func (q *Queries) CreateMessage(ctx context.Context, arg CreateMessageParams) (pgtype.Text, error) {
	row := q.db.QueryRow(ctx, createMessage,
		arg.ID,
		arg.ConnectionID,
		arg.MessageContent,
		arg.UserID,
	)
	var id pgtype.Text
	err := row.Scan(&id)
	return id, err
}

const createUser = `-- name: CreateUser :execrows

INSERT INTO
    "users"(
        id,
        first_name,
        last_name,
        email,
        password,
        image_url,
        description
    )
VALUES ($1, $2, $3, $4, $5, $6, $7)
`

type CreateUserParams struct {
	ID          string      `db:"id"`
	FirstName   string      `db:"first_name"`
	LastName    string      `db:"last_name"`
	Email       string      `db:"email"`
	Password    string      `db:"password"`
	ImageUrl    pgtype.Text `db:"image_url"`
	Description pgtype.Text `db:"description"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (int64, error) {
	result, err := q.db.Exec(ctx, createUser,
		arg.ID,
		arg.FirstName,
		arg.LastName,
		arg.Email,
		arg.Password,
		arg.ImageUrl,
		arg.Description,
	)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected(), nil
}

const createUserConnectionWith = `-- name: CreateUserConnectionWith :one

INSERT INTO
    "user_connections" (user_1, user_2, id)
VALUES ($1, $2, $3) RETURNING id,
    status
`

type CreateUserConnectionWithParams struct {
	User1 string `db:"user_1"`
	User2 string `db:"user_2"`
	ID    string `db:"id"`
}

type CreateUserConnectionWithRow struct {
	ID     string     `db:"id"`
	Status NullStatus `db:"status"`
}

func (q *Queries) CreateUserConnectionWith(ctx context.Context, arg CreateUserConnectionWithParams) (*CreateUserConnectionWithRow, error) {
	row := q.db.QueryRow(ctx, createUserConnectionWith, arg.User1, arg.User2, arg.ID)
	var i CreateUserConnectionWithRow
	err := row.Scan(&i.ID, &i.Status)
	return &i, err
}
