package controllers

import (
	"context"
	"fmt"
	"io"

	"github.com/labstack/echo/v4"
	"github.com/unedtamps/chat-app/api/templates"
)

func (c *Controlers) FindUser(e echo.Context) error {
	a := templates.Hello("ahho")
	var w io.Writer
	b := []int{1,3,4,}
	fmt.Println(b)
	a.Render(context.Background(),w)
	err := e.Render(200,"hi",w)
	if err != nil {
		return err 
	}
	return nil
}