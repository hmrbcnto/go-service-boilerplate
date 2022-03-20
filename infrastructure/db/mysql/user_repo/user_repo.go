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

func (ur *userRepo) CreateUser(user *entities.User) (*entities.User, error) {
	result := ur.db.Create(user)

	if result.Error != nil {
		return nil, result.Error
	}

	return user, result.Error
}
