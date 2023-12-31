package responses

import (
	"fmt"

	userModel "github.com/Surdy-A/amis_portal/internal/modules/user/models"
)

type User struct {
	ID        uint   `json:"id"`
	Image     string `json:"image"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Email     string `json:"email"`
}

type Users struct {
	Data []User
}

func ToUser(user userModel.User) User {
	return User{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.Username,
		Email:     user.Email,
		Image:     fmt.Sprintf("https://ui-avatars.com/api/?name=%s", user.Username),
	}
}
