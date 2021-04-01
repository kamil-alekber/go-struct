package auth

type AuthService interface {
	Migrations() []byte
}

type service struct {
	Repository AuthRepository
}

func NewAuthService(repository AuthRepository) AuthService {
	return &service{Repository: repository}
}

func (s *service) Migrations() []byte {
	res := s.Repository.GetMigrations()
	return res
}
