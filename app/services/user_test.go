package services

import (
	"clerks/domain"
	"clerks/persistence/fakes"
	"errors"
	"testing"
)

func TestUserService_Populate(t *testing.T) {
	ur := &fakes.FakeUserRepository{}
	rr := &fakes.FakeRandomUserRepository{}

	MockUsers := []domain.User{
		{
			Name: domain.Name{
				Title: "Mr",
				First: "John",
				Last:  "Smith",
			},
			Email:       "john@smith.com",
			PhoneNumber: "696969",
			Picture: domain.Picture{
				Thumbnail: "tmb.com",
				Medium:    "mdm.com",
				Large:     "lrg.com",
			},
			RegistrationDate: "2022-05-20T17:22:15.630Z",
		},
		{
			Name: domain.Name{
				Title: "Miss",
				First: "Maria",
				Last:  "Car",
			},
			Email:       "mar@car.com",
			PhoneNumber: "69696",
			Picture: domain.Picture{
				Thumbnail: "https://tmb.com",
				Medium:    "https://mdm.com",
				Large:     "https://lrg.com",
			},
			RegistrationDate: "2021-05-20T17:22:15.630Z",
		},
	}
	rr.FetchUsersReturns(MockUsers, nil)
	ur.AddUsersReturns(nil)

	userService := NewUserService(ur, rr)
	err := userService.Populate()

	if err != nil {
		t.Error(err)
	}
}

func TestUserService_Populate_RandomUserRepository_FAIL(t *testing.T) {
	ur := &fakes.FakeUserRepository{}
	rr := &fakes.FakeRandomUserRepository{}

	MockUsers := []domain.User{}
	rr.FetchUsersReturns(MockUsers, errors.New("response failed"))

	userService := NewUserService(ur, rr)
	err := userService.Populate()

	if err == nil {
		t.Error(err)
	}

}
func TestUserService_Populate_UserRepository_FAIL(t *testing.T) {
	ur := &fakes.FakeUserRepository{}
	rr := &fakes.FakeRandomUserRepository{}

	MockUsers := []domain.User{
		{
			Name: domain.Name{
				Title: "Mr",
				First: "John",
				Last:  "Smith",
			},
			Email:       "john@smith.com",
			PhoneNumber: "696969",
			Picture: domain.Picture{
				Thumbnail: "tmb.com",
				Medium:    "mdm.com",
				Large:     "lrg.com",
			},
			RegistrationDate: "2022-05-20T17:22:15.630Z",
		},
		{
			Name: domain.Name{
				Title: "Miss",
				First: "Maria",
				Last:  "Car",
			},
			Email:       "mar@car.com",
			PhoneNumber: "69696",
			Picture: domain.Picture{
				Thumbnail: "https://tmb.com",
				Medium:    "https://mdm.com",
				Large:     "https://lrg.com",
			},
			RegistrationDate: "2021-05-20T17:22:15.630Z",
		},
	}
	rr.FetchUsersReturns(MockUsers, nil)
	ur.AddUsersReturns(errors.New("failed to create documents"))

	userService := NewUserService(ur, rr)
	err := userService.Populate()

	if err == nil {
		t.Error(err)
	}

}
