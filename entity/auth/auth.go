package auth

type Auth struct {
	Password string `json:"password"`
	UserId   string `json:"user_id"`
}

type Register struct {
	Auth
	Confirm string `json:"confirm"`
}

type Login struct {
	Auth
	Email string `json:"email"`
}

// reflect.TypeOf(user)
