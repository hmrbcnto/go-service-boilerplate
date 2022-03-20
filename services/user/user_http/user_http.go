package user_http

import (
	"net/http"

	"github.com/go-service-boilerplate/services/user/user_usecase"
)

type UserHTTPHandler interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
}

type user_http_handler struct {
	userUseCase user_usecase.UseCase
}

func NewUserHandler(userUseCase user_usecase.UseCase) UserHTTPHandler {
	return &user_http_handler{
		userUseCase: userUseCase,
	}
}
