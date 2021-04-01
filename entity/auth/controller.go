package auth

import (
	"fmt"
	"go-struct/mux"
	"go-struct/utils"
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

func (c controller) Index(w http.ResponseWriter, req *http.Request) {
	utils.SendFile("/templates/index.html", w)
}

func (c controller) Favicon(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "image/x-icon")
	w.Header().Add("Cache-Control", "max-age=31536000")
	utils.SendFile("/static/favicon.ico", w)
}

func (c controller) Migrations(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, string(c.service.Migrations()))
}

func NewAuthController(service AuthService) mux.Router {
	authController := controller{service}
	var r mux.Router

	r.Get("/", authController.Index)
	r.Get("/favicon.ico", authController.Favicon)
	r.Get("/register", authController.Register)
	r.Get("/migrations", authController.Migrations)
	r.Get("/migrations/:id", authController.Migrations)

	return r
}
