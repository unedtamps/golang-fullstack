package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/unedtamps/golang-fullstack/public/web/pages"
	"github.com/unedtamps/golang-fullstack/src/utils"
)

func (c *Controlers) Home(e echo.Context) error {
	tmp := pages.Home()
	return utils.Render(tmp, e)
}
