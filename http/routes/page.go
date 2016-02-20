package routes

import (
	"github.com/cispell/cispell/http/controllers"
	"github.com/labstack/echo"
)

func initPage(e *echo.Echo) {
	e.Get("/", controllers.GetHome)
	e.Get("/terms", controllers.GetTerms)
	e.Get("/privacy", controllers.GetPrivacy)
	e.Get("/about", controllers.GetAbout)
	e.Get("/docs", controllers.GetDocs)
}
