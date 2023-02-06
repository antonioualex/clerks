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

	if limit == 0 {
		limit = 10
	}

	if limit < 0 || limit > 100 {
		return []domain.User{}, 0, 0, errors.New("invalid limit")
	}

	if offset < 0 {
		return []domain.User{}, 0, 0, errors.New("invalid offset")
	}

	aql := `		
		FOR user IN users 
			FILTER @email == '' || lower(user.email) == lower(@email) 
			SORT DATE_TIMESTAMP(user.registered_date) desc 
			LIMIT @input_offset, @limit 
			RETURN user
	`

	bindVars := map[string]interface{}{
		"email":        email,
		"limit":        limit,
		"input_offset": offset * limit,
	}

	ctx := context.Background()
	ctx = driver.WithQueryCount(ctx)

	cursor, err := r.db.Query(ctx, aql, bindVars)
	if err != nil {
		return []domain.User{}, 0, 0, err
	}
	defer cursor.Close()

	if cursor.Count() == 0 {
		return []domain.User{}, 0, 0, errors.New("no more documents")
	}

	var userList []domain.User

	for cursor.HasMore() {
		var user domain.User
		_, readDocumentErr := cursor.ReadDocument(nil, &user)

		if driver.IsNoMoreDocuments(readDocumentErr) {
			break
		} else if readDocumentErr != nil {
			return []domain.User{}, 0, 0, readDocumentErr
		}

		userList = append(userList, user)
	}

	endingBefore := 0
	if offset > 1 {
		endingBefore = offset - 1
	}

	startingAfter := offset + 1

	return userList, startingAfter, endingBefore, nil
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
