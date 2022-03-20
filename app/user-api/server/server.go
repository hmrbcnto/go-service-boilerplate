package server

import (
	"net/http"
	"time"

	"github.com/go-service-boilerplate/infrastructure/db/mysql/user_repo"
	user_handler "github.com/go-service-boilerplate/services/user/user_http"
	"github.com/go-service-boilerplate/services/user/user_usecase"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Server interface {
	ListenAndServe(port string) error
}

type server struct {
	db       *gorm.DB
	mux      *mux.Router
	handlers user_handler.UserHTTPHandler
}

func New(db *gorm.DB) Server {
	mux := mux.NewRouter()
	return &server{
		db:  db,
		mux: mux,
	}
}
func (srv *server) ListenAndServe(port string) error {
	userRepo := user_repo.NewRepo(srv.db)
	userUseCase := user_usecase.New(userRepo)
	srv.handlers = user_handler.NewUserHandler(userUseCase)

	srv.initRoutes()

	httpServer := http.Server{
		Addr:         port,
		Handler:      srv.mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	return httpServer.ListenAndServe()
}
