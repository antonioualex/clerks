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

type OffsetResponse struct {
	StartingAfter int `json:"starting_after"`
	EndingBefore  int `json:"ending_before"`
}

type ClerksResponse struct {
	OffsetResponse `json:"Offset"`
	Results        []domain.User `json:"results"`
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

	offsetStr := request.URL.Query().Get("offset")
	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		offset = 0
	}

	users, startingAfter, endingBefore, err := uh.us.GetUsers(email, limit, offset)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte(err.Error()))
		return
	}

	clerksResp := ClerksResponse{
		OffsetResponse: OffsetResponse{
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
