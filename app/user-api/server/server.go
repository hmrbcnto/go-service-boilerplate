package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	serverError := make(chan error, 1)

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

	go func() {
		log.Printf("server starting at port: %v", port)
		serverError <- httpServer.ListenAndServe()
	}()

	select {
	case <-shutdown:
		log.Println("error")
		return fmt.Errorf("server error")
	case srvErr := <-serverError:
		log.Println("starting graceful server shutdown")
		defer log.Println("server has shutdown successfully")

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := httpServer.Shutdown(ctx); err != nil {
			log.Println("server graceful shutdown is not successful")
			httpServer.Close()
			return err
		}

		<-ctx.Done()
		return srvErr
	}
}
