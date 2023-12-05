// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0

package repository

import (
	"database/sql/driver"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
)

type Status string

const (
	StatusActive Status = "active"
	StatusBlock  Status = "block"
	StatusMute   Status = "mute"
)

func (e *Status) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Status(s)
	case string:
		*e = Status(s)
	default:
		return fmt.Errorf("unsupported scan type for Status: %T", src)
	}
	return nil
}

type NullStatus struct {
	Status Status
	Valid  bool // Valid is true if Status is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullStatus) Scan(value interface{}) error {
	if value == nil {
		ns.Status, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.Status.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullStatus) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.Status), nil
}

func (e Status) Valid() bool {
	switch e {
	case StatusActive,
		StatusBlock,
		StatusMute:
		return true
	}
	return false
}

type Attachment struct {
	CreatedAt     pgtype.Timestamptz `db:"created_at"`
	UpdatedAt     pgtype.Timestamptz `db:"updated_at"`
	IDMessages    string             `db:"id_messages"`
	AttachmentUrl string             `db:"attachment_url"`
	ID            string             `db:"id"`
	ConnectionID  string             `db:"connection_id"`
}

type Group struct {
	ID          string             `db:"id"`
	Name        string             `db:"name"`
	Description pgtype.Text        `db:"description"`
	CreatedAt   pgtype.Timestamptz `db:"created_at"`
	UpdatedAt   pgtype.Timestamptz `db:"updated_at"`
}

type GroupAttachment struct {
	ID            string             `db:"id"`
	AttachmentUrl string             `db:"attachment_url"`
	CreatedAt     pgtype.Timestamptz `db:"created_at"`
	UpdatedAt     pgtype.Timestamptz `db:"updated_at"`
	GroupID       string             `db:"group_id"`
	IDMessage     string             `db:"id_message"`
}

type GroupMessage struct {
	ID             pgtype.Text        `db:"id"`
	UserID         string             `db:"user_id"`
	GroupID        string             `db:"group_id"`
	MessageContent string             `db:"message_content"`
	CreatedAt      pgtype.Timestamptz `db:"created_at"`
	UpdatedAt      pgtype.Timestamptz `db:"updated_at"`
}

type GroupParticipant struct {
	GroupID string `db:"group_id"`
	UserID  string `db:"user_id"`
}

type Message struct {
	ID             pgtype.Text        `db:"id"`
	ConnectionID   string             `db:"connection_id"`
	CreatedAt      pgtype.Timestamptz `db:"created_at"`
	UpdatedAt      pgtype.Timestamptz `db:"updated_at"`
	MessageContent pgtype.Text        `db:"message_content"`
	UserID         string             `db:"user_id"`
}

type User struct {
	ID          string             `db:"id"`
	FirstName   string             `db:"first_name"`
	LastName    string             `db:"last_name"`
	Email       string             `db:"email"`
	Password    string             `db:"password"`
	CreatedAt   pgtype.Timestamptz `db:"created_at"`
	UpdatedAt   pgtype.Timestamptz `db:"updated_at"`
	ImageUrl    pgtype.Text        `db:"image_url"`
	Description pgtype.Text        `db:"description"`
}

type UserConnection struct {
	User1     string             `db:"user_1"`
	User2     string             `db:"user_2"`
	ID        string             `db:"id"`
	CreatedAt pgtype.Timestamptz `db:"created_at"`
	UpdatedAt pgtype.Timestamptz `db:"updated_at"`
	Status    NullStatus         `db:"status"`
}