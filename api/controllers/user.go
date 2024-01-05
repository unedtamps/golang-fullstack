package controllers

import (
	"bytes"
	"context"
	"fmt"
	"io"

	"github.com/labstack/echo/v4"
	"github.com/unedtamps/chat-app/api/templates"
)

func (c *Controlers) FindUser(e echo.Context) error {
	a := templates.Hello("ahho")
	var w io.Writer
	b := []int{1, 3, 4}
	fmt.Println(b)
	a.Render(context.Background(), w)
	err := e.Render(200, "hi", w)
	if err != nil {
		return err
	}
	return nil
}

func (c *Controlers) Home(e echo.Context) error {

	//a := struct {
	//	Name string `json:"name"`
	//	Bar  string `json:"bar"`
	//}{
	//	Name: "bujang lah",
	//	Bar:  "foo",
	//}

	buff := new(bytes.Buffer)

	tmp := templates.Hello("unedo")

	tmp.Render(context.Background(), buff)

	fmt.Println(buff)

	return e.Render(200, "", buff)
}

