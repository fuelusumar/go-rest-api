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
		router.Route{Path: "/edit/", Controller: controller.PageEditHandler},
		router.Route{Path: "/save/", Controller: controller.PageSaveHandler},
		router.Route{Path: "/views/", Controller: controller.PageViewHandler},
	)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
