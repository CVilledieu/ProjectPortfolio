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
	header Header
	pro    Project
}

type Header struct {
	welcome string
}

func getHomePage() *Page {
	return &Page{
		project: GetProject(),
	}
}

func getHeader() Header {
	return Header{
		welcome: "Welcome to my back-end dev portfolio!",
	}
}

type Project struct {
	title       string
	description string
}

func createProject(name, description string) Project {
	return Project{
		title:       name,
		description: description,
	}
}

func GetProject() Project {
	names := "Password Generator"
	descriptions := "A CLI tool to create passwords. Takes in a length between 8-12 and outputs a random series of Uppercase, lowercase, numbers, and symbols"
	return createProject(names, descriptions)
}
