package user

import (
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

type UserRepository interface {
	GetUser(id string) User
	ListUsers([]string) []User
	InsertUser(u *User) bool
	UpdateUser(u *User) bool
	DeleteUser(u *User) bool
}

type repository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) UserRepository {
	return &repository{db: db}
}

func (r *repository) GetUser(id string) User {
	fmt.Println("Returning one user")
	// getting responce from db
	// r.db.Exec("SELECT * FROM users WHERE id={id}", id)

	u := User{FirstName: "Kam", Email: "Kam@gmail.com", LastName: "username", ID: "123"}

	return u
}

func (r *repository) ListUsers([]string) []User {
	fmt.Println("Returning many users")
	// getting repsonce from the db
	u := User{FirstName: "Kam", Email: "Kam@gmail.com", LastName: "username", ID: "123"}
	res := []User{u}
	return res
}

func (r *repository) InsertUser(u *User) bool {
	fmt.Println("Inserting new user")
	return true
}

func (r *repository) UpdateUser(u *User) bool {
	fmt.Println("Updating user")
	return true
}

func (r *repository) DeleteUser(u *User) bool {
	fmt.Println("deleting user")
	return true
}
