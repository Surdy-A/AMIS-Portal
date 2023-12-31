package services

import (
	"errors"

	userModel "github.com/Surdy-A/amis_portal/internal/modules/user/models"
	UserRepository "github.com/Surdy-A/amis_portal/internal/modules/user/repositories"
	"github.com/Surdy-A/amis_portal/internal/modules/user/requests/auth"
	UserResponse "github.com/Surdy-A/amis_portal/internal/modules/user/responses"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepository UserRepository.UserRepositoryInterface
}

func New() *UserService {
	return &UserService{
		userRepository: UserRepository.New(),
	}
}

func (userService *UserService) Create(request auth.RegisterRequest) (UserResponse.User, error) {
	var response UserResponse.User
	var user userModel.User

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), 12)
	if err != nil {
		return response, errors.New("error hashing the password")
	}

	user.FirstName = request.FirstName
	user.LastName = request.LastName
	user.Username = request.Username
	user.Email = request.Email
	user.Password = string(hashedPassword)

	newUser := userService.userRepository.Create(user)

	if newUser.ID == 0 {
		return response, errors.New("error on creating the user")
	}

	return UserResponse.ToUser(newUser), nil
}

func (userService *UserService) CheckUserExists(email string) bool {
	user := userService.userRepository.FindByEmail(email)

	if user.ID != 0 {
		return true
	}

	return false
}

func (userService *UserService) HandleUserLogin(request auth.LoginRequest) (UserResponse.User, error) {
	var response UserResponse.User
	existsUser := userService.userRepository.FindByEmail(request.Email)

	if existsUser.ID == 0 {
		return response, errors.New("invalid credentials")
	}

	err := bcrypt.CompareHashAndPassword([]byte(existsUser.Password), []byte(request.Password))
	if err != nil {
		return response, errors.New("invalid credentials")
	}

	return UserResponse.ToUser(existsUser), nil
}
