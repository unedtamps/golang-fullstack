package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/unedtamps/golang-fullstack/src/controllers"
)

func LoadRoute(c *controllers.Controlers) *echo.Echo {
	e := echo.New()

	e.Static("/", "public/assets")

	// serve home
	e.GET("/", c.Home)

	return e
}
