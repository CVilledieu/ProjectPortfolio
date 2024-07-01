package main

import (
	"html/template"
	"io"
	"net/http"

	proj "Portfolio/project"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.Static("/static", "assets")
	e.Renderer = newTemplate()
	e.GET("/", home)

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

func home(c echo.Context) error {
	page := getHomePage()
	return c.Render(http.StatusOK, "index", page)
}

type Page struct {
	header   *Header
	projects []*proj.Project
}

type Header struct {
	welcome string
}

func getHomePage() *Page {
	return &Page{
		header:   getHeader(),
		projects: proj.GetListOfProjects(),
	}
}

func getHeader() *Header {
	return &Header{
		welcome: "Welcome to my back-end dev portfolio!",
	}
}
