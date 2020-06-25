package router

import "net/http"

//Route route data structure
type Route struct {
	Path       string
	Controller func(http.ResponseWriter, *http.Request)
}

//RegisterRoutes iterates a routes slices and assigns its matching controller
func RegisterRoutes(routes ...Route) {
	for _, route := range routes {
		http.HandleFunc(route.Path, route.Controller)
	}
}

func NewRoute(path string, controller func(w http.ResponseWriter, r *http.Request)) Route {
	return Route{
		Path:       path,
		Controller: controller,
	}
}
