package routes

import (
	"github.com/cispell/cispell/http/controllers"
	"github.com/labstack/echo"
)

func initUser(e *echo.Echo) {
	e.Get("/profile", controllers.GetProfile)
	e.Get("/repos", controllers.GetRepos)
}
