package http

import (
	"github.com/cispell/cispell/http/routes"
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
	"html/template"
	"io"
)

type (
	Server struct {
		Server *echo.Echo
		Router *echo.Echo
	}

	// Template provides HTML template rendering
	Template struct {
		templates *template.Template
	}
)

var App *Server

// Render HTML
func (t *Template) Render(w io.Writer, name string, data interface{}) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func NewServer() {
	App = &Server{}
	App.Server = echo.New()
}

func InitServer() {
	// Echo instance
	e := App.Server

	// Enables debug mode.
	e.Debug()

	// Middleware
	e.Use(mw.Logger())
	e.Use(mw.Recover())
	e.Use(mw.Gzip())

	// Serve favicon
	e.Favicon("public/favicon.ico")

	// Serve static files
	e.Static("/public", "public")

	//-----------
	// Templates
	//-----------

	t := &Template{
		// Cached templates
		templates: template.Must(template.ParseGlob("resources/views/*.html")),
	}
	e.SetRenderer(t)

	// Routes
	routes.InitRoutes(e)
	App.Router = e
}

func StartServer() {
	App.Server.Run(":80")
}
