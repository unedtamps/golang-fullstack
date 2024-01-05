package service

import d "github.com/unedtamps/golang-fullstack/internal/repository"

type ServiceInterface struct {
	User userI
	Post postI
}
type userI interface {
	SearchUser(string) (*d.Attachment, error)
}
type postI interface {
	SearchPost(string) error
}

func NewService(q d.Querier) *ServiceInterface {
	return &ServiceInterface{
		User: &user_service{q},
	}
}

