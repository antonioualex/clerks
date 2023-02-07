package users

import (
	"clerks/app/fakes"
	"clerks/domain"
	"encoding/json"
	"errors"
	"github.com/google/go-cmp/cmp"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserHandler_PopulateService_StatusOK(t *testing.T) {
	srv := &fakes.FakeUserService{}

	srv.PopulateReturns(nil)

	myHandler := UserHandler{srv}
	ts := httptest.NewServer(http.HandlerFunc(myHandler.Populate))

	resp, err := http.Get(ts.URL)
	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Status code is not 201, but is %v", resp.StatusCode)
	}

}

func TestUserHandler_PopulateService_BadRequest(t *testing.T) {

	srv := &fakes.FakeUserService{}
	srv.PopulateReturns(errors.New("failed to fetch users"))
	myHandler := UserHandler{srv}

	// Creating a fake server
	ts := httptest.NewServer(http.HandlerFunc(myHandler.Populate))
	defer ts.Close()

	res, err := http.Post(ts.URL, "application/json", nil)
	if err != nil {
		t.Error(err)
	}

	err = res.Body.Close()
	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected %v but got %v", http.StatusBadRequest, res.StatusCode)
	}
}

func TestUserHandler_GetUsersService_StatusOK(t *testing.T) {
	srv := &fakes.FakeUserService{}

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
	srv.GetUsersReturns(MockUsers, startingAfter, endingBefore, nil)

	myHandler := UserHandler{srv}
	ts := httptest.NewServer(http.HandlerFunc(myHandler.Clerks))

	resp, err := http.Get(ts.URL)
	if err != nil {
		t.Error(err)
	}

	resUser, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status code is not 201, but is %v", resp.StatusCode)
	}

	clerksResp := &ClerksResponse{}
	err = json.Unmarshal(resUser, &clerksResp)
	if err != nil {
		t.Error(err)
	}

	if clerksResp.PageResponse.EndingBefore != endingBefore {
		t.Error("invalid EndingBefore response")
	}

	if clerksResp.PageResponse.StartingAfter != startingAfter {
		t.Error("invalid StartingAfter response")
	}

	if !cmp.Equal(clerksResp.Results, MockUsers) {
		t.Error("users do not match")
	}

}

func TestUserHandler_GetUsersService_BadRequest(t *testing.T) {

	srv := &fakes.FakeUserService{}
	srv.GetUsersReturns([]domain.User{}, 0, 0, errors.New("failed to fetch users"))
	myHandler := UserHandler{srv}

	// Creating a fake server
	ts := httptest.NewServer(http.HandlerFunc(myHandler.Clerks))
	defer ts.Close()

	res, err := http.Post(ts.URL, "application/json", nil)
	if err != nil {
		t.Error(err)
	}

	err = res.Body.Close()
	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected %v but got %v", http.StatusBadRequest, res.StatusCode)
	}
}
