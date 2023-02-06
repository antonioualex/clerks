package services

import "clerks/domain"

type UserService struct {
	ur domain.UserRepository
	rr domain.RandomUserRepository
}

func (s UserService) Populate() error {
	return nil
}

func (s UserService) GetUsers(email string, limit, offset int) ([]domain.User, int, int, error) {
	return []domain.User{}, 0, 0, nil
}
