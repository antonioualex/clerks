package random_user

import "clerks/domain"

type RandomUserRepository struct {
	URL string
}

func (r RandomUserRepository) FetchUsers() ([]domain.User, error) {
	return []domain.User{}, nil
}

func NewRandomUserRepository(URL string) (*RandomUserRepository, error) {
	return &RandomUserRepository{URL: URL}, nil
}
