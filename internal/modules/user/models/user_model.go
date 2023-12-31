package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string `gorm:"varchar:191" json:"first_name"`
	LastName  string `gorm:"varchar:191" json:"last_name"`
	Username  string `gorm:"varchar:191" json:"username"`
	Email     string `gorm:"varchar:191" json:"email"`
	Password  string `gorm:"varchar:191" json:"password"`
}
