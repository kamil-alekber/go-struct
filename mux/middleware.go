package mux

import "net/http"

type Middleware struct {
	Method  string
	Path    string
	Include bool
	Handler func(http.ResponseWriter, *http.Request)
}
