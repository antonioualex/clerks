package services

import (
	"clerks/domain"
	"errors"
	"fmt"
	"log"
	"sync"
)

type UserService struct {
	ur domain.UserRepository
	rr domain.RandomUserRepository
}

func (s UserService) Populate() error {
	randomUsers, err := s.rr.FetchUsers()
	if err != nil {
		return err
	}

	chunkSize := 100
	numChunks := len(randomUsers) / chunkSize
	if len(randomUsers)%chunkSize != 0 {
		numChunks++
	}

	var wg sync.WaitGroup
	wg.Add(numChunks)
	errCh := make(chan error, numChunks)

	for i := 0; i < numChunks; i++ {
		go func(i int) {
			start := i * chunkSize
			end := (i + 1) * chunkSize
			if end > len(randomUsers) {
				end = len(randomUsers)
			}
			err := s.ur.AddUsers(randomUsers[start:end])
			if err != nil {
				log.Printf("Error storing chunk %d: %v", i, err)
				errCh <- errors.New(fmt.Sprintf("Error storing chunk %d: %v", i, err))
			}
			wg.Done()
		}(i)
	}

	wg.Wait()
	close(errCh)

	numOfFailedChanks := len(errCh)
	if numOfFailedChanks > 0 {
		return errors.New(fmt.Sprintf("failed to store %d chunk of users out of %d", numOfFailedChanks, numChunks))
	}

	return nil
}

func (s UserService) GetUsers(email string, limit, offset int) ([]domain.User, int, int, error) {
	return s.ur.GetUsers(email, limit, offset)
}

func NewUserService(ur domain.UserRepository, rr domain.RandomUserRepository) *UserService {
	return &UserService{ur: ur, rr: rr}
}
