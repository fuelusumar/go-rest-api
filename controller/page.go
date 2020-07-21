package controller

import (
	"fmt"
	"go-rest-api/model"
	"log"
	"net/http"
	"regexp"
	"text/template"
)

const templatesPath = "templates/"

var templates = template.Must(template.ParseFiles(templatesPath+"edit.html", templatesPath+"view.html"))
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
	m := validPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return "", fmt.Errorf("invalid Page Title %s", r.URL.Path)
	}
	return m[2], nil
}

//PageViewHandler handles the view request of a page
func PageViewHandler(w http.ResponseWriter, r *http.Request) {
	title, err := getTitle(w, r)
	if err != nil {
		return
	}

	p, err := model.LoadPage(title)
	if err != nil {
		log.Print("Page not Fount: ", title)
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

// PageEditHandler ...
func PageEditHandler(w http.ResponseWriter, r *http.Request) {
	title, err := getTitle(w, r)
	if err != nil {
		return
	}

	p, err := model.LoadPage(title)
	if err != nil {
		p = &model.Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

// PageSaveHandler ...
func PageSaveHandler(w http.ResponseWriter, r *http.Request) {
	title, err := getTitle(w, r)
	if err != nil {
		return
	}

	body := r.FormValue("body")
	p := &model.Page{Title: title, Body: []byte(body)}

	if err := p.Save(); err != nil {
		throw500Error(w, err)
		return
	}

	http.Redirect(w, r, "/views/"+title, http.StatusFound)
}

// renderTemplate: Render Template otherwise throws internal server error
func renderTemplate(w http.ResponseWriter, tmpl string, p *model.Page) {
	if err := templates.ExecuteTemplate(w, tmpl+".html", p); err != nil {
		throw500Error(w, err)
	}
}

func throw500Error(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
}
