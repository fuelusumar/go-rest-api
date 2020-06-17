package controller

import (
	"fmt"
	"go-rest-api/model"
	"net/http"
)

//PageViewHandler handles the view request of a page
func PageViewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/")+1:]
	p, e := model.LoadPage(title)
	if (e == nil) {
		fmt.Fprintf(w, "<h1>%s</h1>", p.Title)
	}
}