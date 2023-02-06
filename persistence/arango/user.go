package arango

import (
	"clerks/domain"
	"context"
	"errors"
	"fmt"
	"github.com/arangodb/go-driver"
)

type UserRepository struct {
	db driver.Database
	c  string
}

func (r UserRepository) AddUsers(users []domain.User) error {
	col, err := r.db.Collection(nil, r.c)
	if err != nil {
		errMessage := fmt.Sprintf("could not get users collection from database. err: %v", err)
		return errors.New(errMessage)
	}

	_, _, err = col.CreateDocuments(nil, users)
	if err != nil {
		createDocErr := fmt.Sprintf("failed to create user documents. err: %+v", err)
		return errors.New(createDocErr)
	}

	return nil
}

func (r UserRepository) GetUsers(email string, limit, offset int) ([]domain.User, int, int, error) {
	return []domain.User{}, 0, 0, nil
}

func NewUserRepository(db driver.Database, collectionName string) (*UserRepository, error) {
	ctx := context.Background()
	colExists, err := db.CollectionExists(ctx, collectionName)
	if err != nil {
		return nil, err
	}

	if !colExists {
		col, createCollErr := db.CreateCollection(ctx, collectionName, nil)
		if createCollErr != nil {
			return nil, createCollErr
		}

		options := driver.EnsurePersistentIndexOptions{
			Unique: true,
			Name:   "unique_email",
		}
		_, _, err = col.EnsurePersistentIndex(ctx, []string{"email"}, &options)
		if err != nil {
			return nil, err
		}
	}

	return &UserRepository{
		db: db,
		c:  collectionName,
	}, nil
}
