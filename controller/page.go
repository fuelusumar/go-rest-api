package controller

import (
	"go-rest-api/model"
	"net/http"
	"text/template"
)

const templates = "templates/"

//PageViewHandler handles the view request of a page
func PageViewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/views/"):]
	p, err := model.LoadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
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

// renderTemplate: Render Template otherwise throws internal server error
func renderTemplate(w http.ResponseWriter, tmpl string, p *model.Page) {
<<<<<<< HEAD

	t, err := template.ParseFiles("templates/" + tmpl + ".html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
=======
	t, _ := template.ParseFiles(templates + tmpl + ".html")
	t.Execute(w, p)
>>>>>>> 55b995564d4a7c5f392786cbc50cd27fba1e155b
}
