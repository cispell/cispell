package routes

import (
	"github.com/cispell/cispell/http/controllers"
	"github.com/labstack/echo"
)

func initAuth(e *echo.Echo) {
	e.Get("/login", controllers.GetLogin)
	e.Get("/login/github", controllers.GetLoginGithub)
	e.Get("/logout", controllers.GetLogout)
}
