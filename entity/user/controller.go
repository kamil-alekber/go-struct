package user

import (
	"fmt"
	"go-struct/mux"
	"go-struct/utils"
	"io"
	"net/http"
)

type controller struct {
	service UserService
}

func (uc *controller) getUser(w http.ResponseWriter, req *http.Request) {
	id := "123"
	u := uc.service.GetUser(id)

	val, _ := utils.Stringify(u)
	vars := mux.Vars(req)
	fmt.Println(vars["id"], vars["weather"])
	io.WriteString(w, fmt.Sprintf("user: %s \n", val))
}

func (uc *controller) getUsers(w http.ResponseWriter, r *http.Request) {
	ids := []string{"123", "321"}
	userList := uc.service.GetUsers(ids)
	val, _ := utils.Stringify(userList)

	io.WriteString(w, fmt.Sprintf("user: %s \n", val))
}

func NewUserController(service UserService) mux.Router {
	userController := controller{service}
	var r mux.Router
	// r.Get("/api/user", userController.getUser)
	r.Get("/api/user/:id/after/:weather", userController.getUser)
	r.Get("/api/users", userController.getUsers)

	return r
}
