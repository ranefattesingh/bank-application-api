package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

const (
	baseURL = "/api/v1"
)

type Route struct {
	Method      string
	Path        string
	HandlerFunc http.HandlerFunc
}

type Router interface {
	Routes() map[string]Route
}

func NewRouter(routers ...Router) *mux.Router {
	mux := mux.NewRouter()

	v1Router := mux.PathPrefix(baseURL).Subrouter()
	for _, router := range routers {
		for _, route := range router.Routes() {
			v1Router.Path(route.Path).Methods(route.Method).HandlerFunc(route.HandlerFunc)
		}
	}

	return v1Router
}
