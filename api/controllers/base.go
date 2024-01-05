package controllers

import (
	"github.com/unedtamps/chat-app/api/service"
)

type Controlers struct {
	*service.ServiceInterface
}

func NewController(s *service.ServiceInterface) *Controlers {
	return &Controlers{s}
}

