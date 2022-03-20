package user_usecase

import (
	"github.com/go-service-boilerplate/entities"
	"github.com/go-service-boilerplate/infrastructure/db/mysql/user_repo"
)

type UseCase interface {
	CreateUser(user *entities.User) (*entities.User, error)
}

type user_usecase struct {
	userRepo user_repo.UserRepo
}

// usecase handles all the business logic
// Arguments for initializing usecase is not limited to DB Repository
// You can add other dependencies such as third-party application (ex. email-sender)
func New(userRepo user_repo.UserRepo) UseCase {
	return &user_usecase{
		userRepo: userRepo,
	}
}

func (userUC *user_usecase) CreateUser(user *entities.User) (*entities.User, error) {
	return userUC.userRepo.CreateUser(user)
}
