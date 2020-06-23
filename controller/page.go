package controller

import (
	"go-rest-api/model"
	"net/http"
	"text/template"
)

//PageViewHandler handles the view request of a page
func PageViewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/views/"):]
	p, _ := model.LoadPage(title)
	renderTemplate(w, "view", p)
}

// PageEditHandler ...
func PageEditHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := model.LoadPage(title)
	if err != nil {
		p = &model.Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

// PageSaveHandler ...
func PageSaveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/save/"):]
	body := r.FormValue("body")
	p := &model.Page{Title: title, Body: []byte(body)}
	p.Save()
	http.Redirect(w, r, "/views/"+title, http.StatusFound)
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *model.Page) {
	t, _ := template.ParseFiles("templates/" + tmpl + ".html")
	t.Execute(w, p)
}
