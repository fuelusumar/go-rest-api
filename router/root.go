package router

import (
	"log"
	"net/http"
	"regexp"
)

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			log.Printf("invalid Page Title %s", r.URL.Path)
			return
		}
		fn(w, r, m[2])
	}
}

//Route route data structure
type Route struct {
	Path       string
	Controller func(http.ResponseWriter, *http.Request, string)
}

//RegisterRoutes iterates a routes slices and assigns its matching controller
func RegisterRoutes(routes ...Route) {
	for _, route := range routes {
		http.HandleFunc(route.Path, makeHandler(route.Controller))
	}
}

func NewRoute(path string, controller func(http.ResponseWriter, *http.Request, string)) Route {
	return Route{
		Path:       path,
		Controller: controller,
	}
}
