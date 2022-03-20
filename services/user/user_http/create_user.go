package user_http

import (
	"encoding/json"
	"net/http"

	"github.com/go-service-boilerplate/entities"
)

func (userHandler *user_http_handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	// insert typical http requesting parsing blablabla
	w.Header().Set("Content-Type", "application/json")
	jsonWriter := json.NewEncoder(w)
	user, err := userHandler.userUseCase.CreateUser(&entities.User{})

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		jsonWriter.Encode(err)
	}

	w.WriteHeader(http.StatusOK)
	jsonWriter.Encode(user)
}
