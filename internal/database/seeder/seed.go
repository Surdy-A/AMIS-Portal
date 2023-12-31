package seeder

import (
	"log"
	"time"

	schoolModel "github.com/Surdy-A/amis_portal/internal/modules/school/models"
	userModel "github.com/Surdy-A/amis_portal/internal/modules/user/models"
	"github.com/Surdy-A/amis_portal/pkg/database"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Seed() {
	db := database.Connection()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("secret"), 12)
	if err != nil {
		log.Fatal("hash password error")
		return
	}
	user := userModel.User{
		Model:     gorm.Model{},
		FirstName: "Ajayi",
		LastName:  "Sodiq",
		Username:  "Surdyhey",
		Email:     "sodiq@email.com",
		Password:  string(hashedPassword),
	}
	db.Create(&user) // pass pointer of data to Create

	log.Printf("User created successfully with email adreess %s \n", user.Email)

	for i := 1; i <= 10; i++ {
		school := schoolModel.School{
			Model: gorm.Model{
				ID:        1,
				CreatedAt: time.Now(),
				UpdatedAt: time.Time{},
				DeletedAt: gorm.DeletedAt{},
			},
			//	SchoolName:              "",
		}
		db.Create(&school) // pass pointer of data to Create

		log.Printf("Article created successfully with title %s \n", school)
	}

	for i := 1; i <= 10; i++ {
		school := schoolModel.Proprietor{
			Model: gorm.Model{
				ID:        1,
				CreatedAt: time.Now(),
				UpdatedAt: time.Time{},
				DeletedAt: gorm.DeletedAt{},
			},
			Name:          "Ajayi Sodeeq A.",
			Qualification: "BSc",
			Phone: "08039749738",
			Occupation: "Software Engineer",
			Address: "Abeokuta",
			SchoolID: 59,
		}
		db.Create(&school) // pass pointer of data to Create
	}

	log.Println("Seeder done ..")
}
