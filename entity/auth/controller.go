package auth

import (
	"fmt"
	"go-struct/mux"
	"io"
	"net/http"
)

type controller struct {
	service AuthService
}

func (c controller) Register(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	fmt.Println(vars["id"], vars["weather"])
	io.WriteString(w, "implement register route")
}

func (c controller) Migrations(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, string(c.service.Migrations()))
}

func NewAuthController(service AuthService) mux.Router {
	authController := controller{service}
	var r mux.Router

	r.Get("/register", authController.Register)
	r.Get("/migrations", authController.Migrations)
	r.Get("/migrations/:id", authController.Migrations)

	return r
}
