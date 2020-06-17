package main

import (
	"fmt"
	"go-rest-api/controller"
	"go-rest-api/router"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, %s!", r.URL.Path[1:])
}

func main() {
	router.RegisterRoutes(
		router.Route{"/mike", handler},
		router.Route{"/wando", handler},
		router.Route{"/marico", handler},
		router.Route{"/views/", controller.PageViewHandler},
	)
	log.Fatal(http.ListenAndServe(":8080", nil))
}


