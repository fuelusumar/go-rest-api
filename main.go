package main

import (
	"fmt"
	c "go-rest-api/controller"
	r "go-rest-api/router"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, %s!", r.URL.Path[1:])
}

func main() {
	r.RegisterRoutes(
		r.NewRoute("/edit/", c.PageEditHandler),
		r.NewRoute("/save/", c.PageSaveHandler),
		r.NewRoute("/views/", c.PageViewHandler),
	)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
