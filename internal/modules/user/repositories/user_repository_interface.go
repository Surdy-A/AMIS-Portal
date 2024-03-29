package repositories

import (
	userModel "github.com/Surdy-A/amis_portal/internal/modules/user/models"
)

type UserRepositoryInterface interface {
	Create(user userModel.User) userModel.User
	FindByEmail(email string) userModel.User
	FindByID(id int) userModel.User
}
