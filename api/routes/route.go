package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/unedtamps/chat-app/api/controllers"
)

func LoadRoute(c *controllers.Controlers) *echo.Echo {
	e := echo.New()

	//user route
	e.GET("/test", c.FindUser)

	// serve home
	e.GET("/testing", c.Home)

	return e
}

