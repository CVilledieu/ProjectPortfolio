package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Template struct {
	templates *template.Template
}

type Page struct {
	Header  Header
	Project []Project
}

type Header struct {
	Welcome string
}

type Project struct {
	Title       string
	Description string
	Learning    string
}

func main() {
	e := echo.New()
	e.Static("/static", "assets")
	e.Renderer = newTemplate()
	e.GET("/", home)

	e.Logger.Fatal(e.Start(":420"))
}

// HTTP GET request for "/"
func home(c echo.Context) error {
	page := getHomePage()
	return c.Render(http.StatusOK, "index", page)
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func newTemplate() *Template {
	return &Template{
		templates: template.Must(template.ParseGlob("assets/views/*.html")),
	}
}

func getHomePage() Page {
	return Page{
		Header:  getHeader(),
		Project: createProjectList(),
	}
}

func getHeader() Header {
	return Header{
		Welcome: "Welcome to my back-end dev portfolio!",
	}
}

func createProject(n, d, l string) Project {
	return Project{
		Title:       n,
		Description: d,
		Learning:    l,
	}
}

func createProjectList() []Project {
	list := []Project{}
	names := []string{"Password Generator", "My own Database"}
	descrip := []string{
		"A CLI tool to create passwords. Takes in a length between 8-12 and outputs a random series of Uppercase, lowercase, numbers, and symbols",
		"A Redis inspired database",
	}
	learn := []string{
		"My focus and motivation was working with the standard reader and writer",
		"My focus was to actively use a more difficult data structure. I had built plenty of the common ones, but rarely needed to use them in my personal projects.",
	}
	for i := range names {
		list = append(list, createProject(names[i], descrip[i], learn[i]))
	}
	return list
}
