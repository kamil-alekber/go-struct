package mux

import (
	"net/http"
)

type route struct {
	// get, post, put
	method string
	// /hello, hello/{id}
	path string

	handler func(http.ResponseWriter, *http.Request)
}

type Router struct {
	routes []route
	Prefix string
}

func (r *Router) Get(path string, handler func(http.ResponseWriter, *http.Request)) {
	path = r.Prefix + path
	newRoute := route{method: http.MethodGet, path: path, handler: handler}
	r.routes = append(r.routes, newRoute)
}

func (r *Router) Post(path string, handler func(http.ResponseWriter, *http.Request)) {
	path = r.Prefix + path
	newRoute := route{method: http.MethodPost, path: path, handler: handler}
	r.routes = append(r.routes, newRoute)
}
