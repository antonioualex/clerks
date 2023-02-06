package users

import (
	"clerks/domain"
	"net/http"
)

type UserHandler struct {
	us domain.UserService
}

func (uh UserHandler) Populate(writer http.ResponseWriter, request *http.Request) {}

func (uh UserHandler) Clerks(writer http.ResponseWriter, request *http.Request) {}

func NewUserHandler(us domain.UserService) UserHandler {
	return UserHandler{us: us}
}
