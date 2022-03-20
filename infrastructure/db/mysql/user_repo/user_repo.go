package user_repo

import (
	"github.com/go-service-boilerplate/entities"
	"gorm.io/gorm"
)

type UserRepo interface {
	CreateUser(user *entities.User) (*entities.User, error)
}

type userRepo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) UserRepo {
	return &userRepo{
		db: db,
	}
}
