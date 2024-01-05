package service

import (
	"context"

	d "github.com/unedtamps/golang-fullstack/internal/repository"
)

type user_service struct {
	d.Querier
}

func (u *user_service) SearchUser(id string) (*d.Attachment, error) {
	a, err := u.FindAttachmentByID(context.Background(), id)
	if err != nil {
		return nil, err
	}

	return a, nil
}

