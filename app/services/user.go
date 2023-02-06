package services

import "clerks/domain"

type UserService struct {
	ur domain.UserRepository
	rr domain.RandomUserRepository
}

func (s UserService) Populate() error {
	randomUsers, err := s.rr.FetchUsers()
	if err != nil {
		return err
	}
	return s.ur.AddUsers(randomUsers)
}

func (s UserService) GetUsers(email string, limit, offset int) ([]domain.User, int, int, error) {
	return []domain.User{}, 0, 0, nil
}

func NewUserService(ur domain.UserRepository, rr domain.RandomUserRepository) *UserService {
	return &UserService{ur: ur, rr: rr}
}
