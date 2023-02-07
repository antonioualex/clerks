package users

import (
	"clerks/domain"
	"encoding/json"
	"net/http"
	"strconv"
)

type UserHandler struct {
	us domain.UserService
}

type PageResponse struct {
	StartingAfter int `json:"starting_after"`
	EndingBefore  int `json:"ending_before"`
}

type ClerksResponse struct {
	PageResponse `json:"page"`
	Results      []domain.User `json:"results"`
}

func (uh UserHandler) Populate(writer http.ResponseWriter, request *http.Request) {
	err := uh.us.Populate()
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte(err.Error()))
		return
	}

	writer.WriteHeader(http.StatusCreated)
	return
}

func (uh UserHandler) Clerks(writer http.ResponseWriter, request *http.Request) {

	email := request.URL.Query().Get("email")

	limitStr := request.URL.Query().Get("limit")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 10
	}

	pageStr := request.URL.Query().Get("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 0
	}

	users, startingAfter, endingBefore, err := uh.us.GetUsers(email, limit, page)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte(err.Error()))
		return
	}

	clerksResp := ClerksResponse{
		PageResponse: PageResponse{
			StartingAfter: startingAfter,
			EndingBefore:  endingBefore,
		},
		Results: users,
	}

	writer.WriteHeader(http.StatusOK)
	result, _ := json.Marshal(clerksResp)
	writer.Write(result)
	return
}

func NewUserHandler(us domain.UserService) UserHandler {
	return UserHandler{us: us}
}
