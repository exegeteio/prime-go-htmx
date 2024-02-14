package server

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"html/template"
	"io"
	"os"
)

func NewEchoServer() *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Renderer = newTemplate()
	return e
}

type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, _ echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func newTemplate() *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
}

func Start(e *echo.Echo) {
	// Start the server.
	// get the PORT variable out of the environment variables
	// if it is not set, default to 3000
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}
