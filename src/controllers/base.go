package controllers

import (
	"github.com/unedtamps/golang-fullstack/src/service"
)

type Controlers struct {
	*service.ServiceInterface
}

func NewController(s *service.ServiceInterface) *Controlers {
	return &Controlers{s}
}
