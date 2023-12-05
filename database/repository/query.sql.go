// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: query.sql

package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const countUserConnections = `-- name: CountUserConnections :one
SELECT COUNT(user_1 = $1 or user_2 = $1) as cout_contacts
FROM "user_connections"
`

func (q *Queries) CountUserConnections(ctx context.Context, user1 string) (int64, error) {
	row := q.db.QueryRow(ctx, countUserConnections, user1)
	var cout_contacts int64
	err := row.Scan(&cout_contacts)
	return cout_contacts, err
}

const findAttachmentByID = `-- name: FindAttachmentByID :one
SELECT created_at, updated_at, id_messages, attachment_url, id, connection_id FROM attachment
WHERE id = $1 LIMIT 1 OFFSET 0
`

func (q *Queries) FindAttachmentByID(ctx context.Context, id string) (*Attachment, error) {
	row := q.db.QueryRow(ctx, findAttachmentByID, id)
	var i Attachment
	err := row.Scan(
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.IDMessages,
		&i.AttachmentUrl,
		&i.ID,
		&i.ConnectionID,
	)
	return &i, err
}

const findMessageByConnectionID = `-- name: FindMessageByConnectionID :many
SELECT id, connection_id, created_at, updated_at, message_content, user_id FROM messages
WHERE connection_id = $1
ORDER BY created_at ASC LIMIT $2 OFFSET $3
`

type FindMessageByConnectionIDParams struct {
	ConnectionID string `db:"connection_id"`
	Limit        int64  `db:"limit"`
	Offset       int64  `db:"offset"`
}

func (q *Queries) FindMessageByConnectionID(ctx context.Context, arg FindMessageByConnectionIDParams) ([]*Message, error) {
	rows, err := q.db.Query(ctx, findMessageByConnectionID, arg.ConnectionID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*Message
	for rows.Next() {
		var i Message
		if err := rows.Scan(
			&i.ID,
			&i.ConnectionID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.MessageContent,
			&i.UserID,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const findMessageByID = `-- name: FindMessageByID :one
SELECT id, connection_id, created_at, updated_at, message_content, user_id FROM messages
WHERE id = $1 LIMIT 1 OFFSET 0
`

func (q *Queries) FindMessageByID(ctx context.Context, id pgtype.Text) (*Message, error) {
	row := q.db.QueryRow(ctx, findMessageByID, id)
	var i Message
	err := row.Scan(
		&i.ID,
		&i.ConnectionID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.MessageContent,
		&i.UserID,
	)
	return &i, err
}

const findUserByID = `-- name: FindUserByID :one
SELECT id, first_name, last_name, email, password, created_at, updated_at, image_url, description FROM "users" WHERE id = $1 LIMIT 1 OFFSET 0
`

func (q *Queries) FindUserByID(ctx context.Context, id string) (*User, error) {
	row := q.db.QueryRow(ctx, findUserByID, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.ImageUrl,
		&i.Description,
	)
	return &i, err
}

const findUserConnectionsByID = `-- name: FindUserConnectionsByID :one
SELECT user_1, user_2, id, created_at, updated_at, status FROM "user_connections" WHERE id = $1 LIMIT 1 OFFSET 0
`

func (q *Queries) FindUserConnectionsByID(ctx context.Context, id string) (*UserConnection, error) {
	row := q.db.QueryRow(ctx, findUserConnectionsByID, id)
	var i UserConnection
	err := row.Scan(
		&i.User1,
		&i.User2,
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Status,
	)
	return &i, err
}

const findUserConnectionsWith = `-- name: FindUserConnectionsWith :one
SELECT user_1, user_2, id, created_at, updated_at, status FROM "user_connections" 
WHERE user_1 = $1 and user_2 = $2 
OR user_1 = $2 and user_2 = $1
LIMIT 1 OFFSET 0
`

type FindUserConnectionsWithParams struct {
	User1 string `db:"user_1"`
	User2 string `db:"user_2"`
}

func (q *Queries) FindUserConnectionsWith(ctx context.Context, arg FindUserConnectionsWithParams) (*UserConnection, error) {
	row := q.db.QueryRow(ctx, findUserConnectionsWith, arg.User1, arg.User2)
	var i UserConnection
	err := row.Scan(
		&i.User1,
		&i.User2,
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Status,
	)
	return &i, err
}

const lisAttachmentByMessageID = `-- name: LisAttachmentByMessageID :many
SELECT created_at, updated_at, id_messages, attachment_url, id, connection_id FROM attachment
WHERE id_messages = $1
`

func (q *Queries) LisAttachmentByMessageID(ctx context.Context, idMessages string) ([]*Attachment, error) {
	rows, err := q.db.Query(ctx, lisAttachmentByMessageID, idMessages)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*Attachment
	for rows.Next() {
		var i Attachment
		if err := rows.Scan(
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.IDMessages,
			&i.AttachmentUrl,
			&i.ID,
			&i.ConnectionID,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listAttachmentByUserCon = `-- name: ListAttachmentByUserCon :many
SELECT created_at, updated_at, id_messages, attachment_url, id, connection_id FROM attachment
WHERE connection_id = $1 LIMIT $2 OFFSET $3
`

type ListAttachmentByUserConParams struct {
	ConnectionID string `db:"connection_id"`
	Limit        int64  `db:"limit"`
	Offset       int64  `db:"offset"`
}

func (q *Queries) ListAttachmentByUserCon(ctx context.Context, arg ListAttachmentByUserConParams) ([]*Attachment, error) {
	rows, err := q.db.Query(ctx, listAttachmentByUserCon, arg.ConnectionID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*Attachment
	for rows.Next() {
		var i Attachment
		if err := rows.Scan(
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.IDMessages,
			&i.AttachmentUrl,
			&i.ID,
			&i.ConnectionID,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listGroupAttachments = `-- name: ListGroupAttachments :many
SELECT id, attachment_url, created_at, updated_at, group_id, id_message FROM group_attachment
WHERE group_id = $1 LIMIT $2 OFFSET $3
`

type ListGroupAttachmentsParams struct {
	GroupID string `db:"group_id"`
	Limit   int64  `db:"limit"`
	Offset  int64  `db:"offset"`
}

func (q *Queries) ListGroupAttachments(ctx context.Context, arg ListGroupAttachmentsParams) ([]*GroupAttachment, error) {
	rows, err := q.db.Query(ctx, listGroupAttachments, arg.GroupID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*GroupAttachment
	for rows.Next() {
		var i GroupAttachment
		if err := rows.Scan(
			&i.ID,
			&i.AttachmentUrl,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.GroupID,
			&i.IDMessage,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listGroupMessages = `-- name: ListGroupMessages :many
SELECT id, user_id, group_id, message_content, created_at, updated_at FROM group_messages
WHERE group_id = $1 LIMIT $2 OFFSET $3
`

type ListGroupMessagesParams struct {
	GroupID string `db:"group_id"`
	Limit   int64  `db:"limit"`
	Offset  int64  `db:"offset"`
}

func (q *Queries) ListGroupMessages(ctx context.Context, arg ListGroupMessagesParams) ([]*GroupMessage, error) {
	rows, err := q.db.Query(ctx, listGroupMessages, arg.GroupID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*GroupMessage
	for rows.Next() {
		var i GroupMessage
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.GroupID,
			&i.MessageContent,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listGroupParticipants = `-- name: ListGroupParticipants :many
SELECT user_id as participants from group_participants
WHERE group_id = $1
`

func (q *Queries) ListGroupParticipants(ctx context.Context, groupID string) ([]string, error) {
	rows, err := q.db.Query(ctx, listGroupParticipants, groupID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var participants string
		if err := rows.Scan(&participants); err != nil {
			return nil, err
		}
		items = append(items, participants)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listGroupsByName = `-- name: ListGroupsByName :many
SELECT id, name, description, created_at, updated_at FROM groups
WHERE name LIKE '%' || $1 || '%'
`

func (q *Queries) ListGroupsByName(ctx context.Context, dollar_1 pgtype.Text) ([]*Group, error) {
	rows, err := q.db.Query(ctx, listGroupsByName, dollar_1)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*Group
	for rows.Next() {
		var i Group
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listMessageAttachments = `-- name: ListMessageAttachments :many
SELECT id, attachment_url, created_at, updated_at, group_id, id_message FROM group_attachment
WHERE id_message = $1
`

func (q *Queries) ListMessageAttachments(ctx context.Context, idMessage string) ([]*GroupAttachment, error) {
	rows, err := q.db.Query(ctx, listMessageAttachments, idMessage)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*GroupAttachment
	for rows.Next() {
		var i GroupAttachment
		if err := rows.Scan(
			&i.ID,
			&i.AttachmentUrl,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.GroupID,
			&i.IDMessage,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listUserConnectWith = `-- name: ListUserConnectWith :many
SELECT user_1 as contacts
FROM "user_connections" WHERE user_2 = $1
UNION
SELECT user_2 as contacts
FROM "user_connections" WHERE user_1 = $1
`

func (q *Queries) ListUserConnectWith(ctx context.Context, dollar_1 pgtype.Text) ([]string, error) {
	rows, err := q.db.Query(ctx, listUserConnectWith, dollar_1)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var contacts string
		if err := rows.Scan(&contacts); err != nil {
			return nil, err
		}
		items = append(items, contacts)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listUserConnectionsID = `-- name: ListUserConnectionsID :many
SELECT id FROM "user_connections" WHERE user_1 = $1 OR user_2 = $1
`

func (q *Queries) ListUserConnectionsID(ctx context.Context, user1 string) ([]string, error) {
	rows, err := q.db.Query(ctx, listUserConnectionsID, user1)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		items = append(items, id)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listUserGroups = `-- name: ListUserGroups :many
SELECT group_id from group_participants
WHERE user_id = $1 LIMIT $2 OFFSET $3
`

type ListUserGroupsParams struct {
	UserID string `db:"user_id"`
	Limit  int64  `db:"limit"`
	Offset int64  `db:"offset"`
}

func (q *Queries) ListUserGroups(ctx context.Context, arg ListUserGroupsParams) ([]string, error) {
	rows, err := q.db.Query(ctx, listUserGroups, arg.UserID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var group_id string
		if err := rows.Scan(&group_id); err != nil {
			return nil, err
		}
		items = append(items, group_id)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listUserMessages = `-- name: ListUserMessages :many
SELECT id, connection_id, created_at, updated_at, message_content, user_id FROM messages
WHERE user_id = $1 LIMIT $2 OFFSET $3
`

type ListUserMessagesParams struct {
	UserID string `db:"user_id"`
	Limit  int64  `db:"limit"`
	Offset int64  `db:"offset"`
}

func (q *Queries) ListUserMessages(ctx context.Context, arg ListUserMessagesParams) ([]*Message, error) {
	rows, err := q.db.Query(ctx, listUserMessages, arg.UserID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*Message
	for rows.Next() {
		var i Message
		if err := rows.Scan(
			&i.ID,
			&i.ConnectionID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.MessageContent,
			&i.UserID,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listUserMessagesByConetent = `-- name: ListUserMessagesByConetent :many
SELECT id, connection_id, created_at, updated_at, message_content, user_id FROM messages
WHERE message_content LIKE '%'|| $1 || '%'
`

func (q *Queries) ListUserMessagesByConetent(ctx context.Context, dollar_1 pgtype.Text) ([]*Message, error) {
	rows, err := q.db.Query(ctx, listUserMessagesByConetent, dollar_1)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*Message
	for rows.Next() {
		var i Message
		if err := rows.Scan(
			&i.ID,
			&i.ConnectionID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.MessageContent,
			&i.UserID,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listUsers = `-- name: ListUsers :many
SELECT id, first_name, last_name, email, password, created_at, updated_at, image_url, description FROM "users" LIMIT $1 OFFSET $2
`

type ListUsersParams struct {
	Limit  int64 `db:"limit"`
	Offset int64 `db:"offset"`
}

func (q *Queries) ListUsers(ctx context.Context, arg ListUsersParams) ([]*User, error) {
	rows, err := q.db.Query(ctx, listUsers, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.FirstName,
			&i.LastName,
			&i.Email,
			&i.Password,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.ImageUrl,
			&i.Description,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}