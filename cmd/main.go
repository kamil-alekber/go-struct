package main

import (
	"context"
	"fmt"
	"go-struct/entity/auth"
	"go-struct/entity/user"
	"go-struct/mux"
	"go-struct/utils"
	"net/http"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

func main() {
	dbpool, err := pgxpool.Connect(context.Background(), "postgresql://user:test@localhost:5432/student")
	if err != nil {
		fmt.Printf("Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()
	// apply migrations
	utils.MigrationUp(dbpool)

	// User Service
	userRepository := user.NewUserRepository(dbpool)
	userService := user.NewUserService(userRepository)
	userRouter := user.NewUserController(userService)

	// Auth Service
	authRepository := auth.NewAuthRepository(dbpool)
	authService := auth.NewAuthService(authRepository)
	authRouter := auth.NewAuthController(authService)

	m := mux.NewMultiplexer()

	whenPathIncludes := mux.Middleware{Include: true, Path: "/api", Handler: func(rw http.ResponseWriter, r *http.Request) {
		fmt.Println("if path includes /api")
	}}

	m.Use(whenPathIncludes)
	m.Register(userRouter, authRouter)
	// should distinguish between middlewares and routes

	// Blog service
	if err := http.ListenAndServe("127.0.0.1:5000", m); err != nil {
		fmt.Printf("Error starting server: %s", err)
		os.Exit(1)
	}
}
