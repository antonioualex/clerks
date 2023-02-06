package random_user

import (
	"clerks/domain"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

type RandomUserRepository struct {
	URL string
}

type RandomUserResponse struct {
	Results []RandomUser `json:"results"`
}

type RandomUser struct {
	Name             RandomName       `json:"name"`
	Email            string           `json:"email"`
	PhoneNumber      string           `json:"phone"`
	Picture          RandomPicture    `json:"picture"`
	RegistrationDate RandomRegistered `json:"registered"`
}

type RandomRegistered struct {
	Date string `json:"date"`
}

type RandomName struct {
	Title string `json:"title"`
	First string `json:"first"`
	Last  string `json:"last"`
}

type RandomPicture struct {
	Thumbnail string `json:"thumbnail"`
	Medium    string `json:"medium"`
	Large     string `json:"large"`
}

func (r RandomUserRepository) FetchUsers() ([]domain.User, error) {
	res, err := http.Get(r.URL)
	if err != nil {
		log.Println(err)
		return []domain.User{}, err
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > http.StatusCreated {
		respErr := fmt.Sprintf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
		return []domain.User{}, errors.New(respErr)
	}
	if err != nil {
		respErr := fmt.Sprintf("failed to read response body, err: %s", err)
		return []domain.User{}, errors.New(respErr)
	}

	response := &RandomUserResponse{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		respErr := fmt.Sprintf("failed to unmarshal response body, err: %+v", err)
		return []domain.User{}, errors.New(respErr)
	}

	var fetchedUsers []domain.User

	for _, randomUser := range response.Results {
		user := domain.User{
			Name:             domain.Name(randomUser.Name),
			Email:            randomUser.Email,
			PhoneNumber:      randomUser.PhoneNumber,
			Picture:          domain.Picture(randomUser.Picture),
			RegistrationDate: randomUser.RegistrationDate.Date,
		}
		fetchedUsers = append(fetchedUsers, user)
	}

	return fetchedUsers, nil
}

func NewRandomUserRepository(URL string) (*RandomUserRepository, error) {
	return &RandomUserRepository{URL: URL}, nil
}
