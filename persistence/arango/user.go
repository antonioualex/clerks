package arango

import "clerks/domain"

type UserRepository struct {
}

func (r UserRepository) AddUsers(users []domain.User) error {
	return nil
}

func (r UserRepository) GetUsers(email string, limit, offset int) ([]domain.User, int, int, error) {
	return []domain.User{}, 0, 0, nil
}

func NewUserRepository() (*UserRepository, error) {
	return &UserRepository{}, nil
}
