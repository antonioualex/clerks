package users

import (
	"clerks/domain"
	"net/http"
)

type UserHandler struct {
	us domain.UserService
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

func (uh UserHandler) Clerks(writer http.ResponseWriter, request *http.Request) {}

func NewUserHandler(us domain.UserService) UserHandler {
	return UserHandler{us: us}
}
