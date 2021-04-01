package user

type UserService interface {
	NewUser(u *User) bool
	GetUser(id string) User
	GetUsers(ids []string) []User
}

type service struct {
	Repository UserRepository
}

func (s *service) NewUser(u *User) bool {
	res := s.Repository.InsertUser(u)

	return res
}

func (s *service) GetUser(id string) User {
	user := s.Repository.GetUser(id)

	return user
}

func (s *service) GetUsers(ids []string) []User {
	users := s.Repository.ListUsers(ids)

	return users
}

func NewUserService(repository UserRepository) UserService {
	return &service{
		Repository: repository,
	}
}
