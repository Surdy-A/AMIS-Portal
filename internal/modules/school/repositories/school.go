package repositories

import (
	"github.com/Surdy-A/amis_portal/internal/modules/school/models"
	"github.com/Surdy-A/amis_portal/pkg/database"
	"gorm.io/gorm"
)

type SchoolRepositoryInterface interface {
	GetSchools(limit int) []models.School
	GetSchool(id int) models.School
	Create(school models.School) models.School
	DeleteSchool(id int) error
}

type SchoolRepository struct {
	DB *gorm.DB
}

func New() *SchoolRepository {
	return &SchoolRepository{
		DB: database.Connection(),
	}
}

func (schoolRepo *SchoolRepository) Create(school models.School) models.School {
	var newSchool models.School

	schoolRepo.DB.Create(&school).Scan(&newSchool)

	return newSchool
}

func (schoolRepo *SchoolRepository) GetSchool(id int) models.School {
	var school models.School

	schoolRepo.DB.First(&school, id)

	return school
}

func (schoolRepo *SchoolRepository) GetSchools(limit int) []models.School {
	var schools []models.School

	schoolRepo.DB.Limit(limit).Joins("User").Find(&schools)

	//schools = append(schools, sch)
	return schools
}

func (schoolRepo *SchoolRepository) DeleteSchool(id int) error {
	var school models.School

	schoolRepo.DB.Unscoped().Delete(&school, id)

	return nil
}
