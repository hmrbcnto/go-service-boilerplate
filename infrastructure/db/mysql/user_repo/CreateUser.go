package user_repo

import "github.com/go-service-boilerplate/entities"

func (ur *userRepo) CreateUser(user *entities.User) (*entities.User, error) {
	result := ur.db.Create(user)

	if result.Error != nil {
		return nil, result.Error
	}

	return user, result.Error
}
