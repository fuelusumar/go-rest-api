package main

import (
	c "go-rest-api/controller"
	r "go-rest-api/router"
	"log"
	"net/http"
)

func main() {
	r.RegisterRoutes(
		r.NewRoute("/edit/", c.PageEditHandler),
		r.NewRoute("/save/", c.PageSaveHandler),
		r.NewRoute("/view/", c.PageViewHandler),
	)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
