package mux

import (
	"fmt"
	"go-struct/utils"
	"io"
	"net/http"
	"strconv"
	"strings"
)

type mux struct {
	routes      []route
	middlewares []Middleware
}

func NewMultiplexer() *mux {
	return &mux{}
}

func (mux *mux) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	method := req.Method
	protocol := req.Proto

	fmt.Printf("%s %s %s \r\n", strings.ToUpper(method), path, protocol)

	for _, middleware := range mux.middlewares {
		pathMatch := path == middleware.Path || path == fmt.Sprintf("%v/", middleware.Path)
		methodMatch := strings.EqualFold(middleware.Method, method)

		if middleware.Handler == nil {
			fmt.Println("One of your middleware does not have Handler()")
		} else if middleware.Path == "" && middleware.Method == "" {
			middleware.Handler(w, req)
		} else if methodMatch && middleware.Path == "" {
			middleware.Handler(w, req)
		} else if pathMatch && middleware.Method == "" {
			middleware.Handler(w, req)
		} else if pathMatch && methodMatch {
			middleware.Handler(w, req)
		} else if middleware.Include && strings.HasPrefix(path, middleware.Path) {
			middleware.Handler(w, req)
		}
	}

	for _, route := range mux.routes {
		pathMatch := path == route.path || path == fmt.Sprintf("%s/", route.path)
		methodMatch := strings.EqualFold(route.method, method)

		if pathMatch && methodMatch {
			route.handler(w, req)
			return
		} else if methodMatch {
			// router => /api/user/:id/after
			// actual => /api/user/1/after
			routerPathChunks := strings.Split(strings.TrimSuffix(route.path, "/"), "/")
			actualPathChunks := strings.Split(strings.TrimSuffix(path, "/"), "/")

			// router => ["api", "user", ":id", "after"]
			// actual => ["api", "user", "123", "after"]

			if len(routerPathChunks) == len(actualPathChunks) {
				hit := false
				dynamicParts := make(map[string]string)
				for i, v := range routerPathChunks {
					if strings.HasPrefix(v, ":") {
						hit = true
						dynamicParts[v[1:]] = actualPathChunks[i]
					}
				}
				str, _ := utils.Stringify(dynamicParts)
				req.Header.Add("vars", str)

				if hit {
					route.handler(w, req)
					return
				}
			}

		}
	}

	default_response := fmt.Sprintf("%s %s is not defined", method, path)
	w.Header().Add("Content-Length", strconv.Itoa(len(default_response)))
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(404)

	io.WriteString(w, default_response)
}

func Vars(req *http.Request) map[string]interface{} {
	vars, _ := utils.Parse([]byte(req.Header.Get("vars")))
	assertedVars, ok := vars.(map[string]interface{})

	if ok {
		return assertedVars
	}
	return map[string]interface{}{}

}

func (mux *mux) Use(middlewares ...Middleware) {
	mux.middlewares = append(mux.middlewares, middlewares...)
}

func (mux *mux) Register(routers ...Router) {
	for _, router := range routers {
		mux.routes = append(mux.routes, router.routes...)
	}
}
