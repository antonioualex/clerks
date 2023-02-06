package users

import (
	"clerks/app/fakes"
	"errors"
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
