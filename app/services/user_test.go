package services

import (
	"clerks/domain"
	"clerks/persistence/fakes"
	"errors"
	"github.com/google/go-cmp/cmp"
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

func TestUserService_GetUsers(t *testing.T) {

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

	endingBefore := 0
	startingAfter := 1

	ur.GetUsersReturns(MockUsers, startingAfter, endingBefore, nil)

	userService := NewUserService(ur, rr)

	emailInput := ""
	limitInput := 2
	offsetInput := 0

	usersResults, startingAfterResult, endingBeforeResult, err := userService.GetUsers(emailInput, limitInput, offsetInput)
	if err != nil {
		t.Error(err)
	}

	if startingAfterResult != startingAfter {
		t.Error("startingAfter result is not valid")
	}
	if endingBeforeResult != endingBefore {
		t.Error("endingBefore result is not valid")
	}

	if !cmp.Equal(usersResults, MockUsers) {
		t.Error("users do not match")
	}
}

func TestUserService_GetUsers_FAIL(t *testing.T) {
	ur := &fakes.FakeUserRepository{}
	rr := &fakes.FakeRandomUserRepository{}

	MockUsers := []domain.User{}

	endingBefore := 0
	startingAfter := 0

	ur.GetUsersReturns(MockUsers, startingAfter, endingBefore, errors.New("invalid limit value"))

	userService := NewUserService(ur, rr)

	emailInput := ""
	limitInput := 0
	offsetInput := 0

	usersResults, startingAfterResult, endingBeforeResult, err := userService.GetUsers(emailInput, limitInput, offsetInput)

	if err == nil {
		t.Error(err)
	}

	if startingAfterResult != startingAfter {
		t.Error("startingAfter result is not valid")
	}
	if endingBeforeResult != endingBefore {
		t.Error("endingBefore result is not valid")
	}

	if !cmp.Equal(usersResults, MockUsers) {
		t.Error("users do not match")
	}
}
