package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.Static("/static", "assets")
	e.Renderer = newTemplate()
	e.GET("/", homeBase)

	e.Logger.Fatal(e.Start(":420"))
}

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func newTemplate() *Template {
	return &Template{
		templates: template.Must(template.ParseGlob("assets/views/*.html")),
	}
}

func homeBase(c echo.Context) error {
	return c.Render(http.StatusOK, "index", nil)
}
