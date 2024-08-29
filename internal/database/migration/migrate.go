package migration

import (
	"fmt"
	"log"

	"github.com/Surdy-A/amis_portal/internal/modules/school/models"
	gradingModels "github.com/Surdy-A/amis_portal/internal/modules/examination/models"
	userModels "github.com/Surdy-A/amis_portal/internal/modules/user/models"
	"github.com/Surdy-A/amis_portal/pkg/database"
)

func Migrate() {
	db := database.Connection()

	err := db.AutoMigrate(&userModels.User{}, &models.School{}, &models.Proprietor{}, 
		&gradingModels.GradingExamination{}, &gradingModels.StudentGradingExamInfo{}, &gradingModels.Blog{})

	if err != nil {
		log.Fatal("Cant migrate")
		return
	}

	fmt.Println("Migration done ..")
}
