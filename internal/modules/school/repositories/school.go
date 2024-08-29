package repositories

import (
	"time"

	"github.com/Surdy-A/amis_portal/internal/modules/school/models"
	"github.com/Surdy-A/amis_portal/pkg/database"
	"gorm.io/gorm"
)

type SchoolRepositoryInterface interface {
	GetSchools(limit int) []models.School
	GetSchool(id int) models.School
	GetSchoolBySchoolCode(schoolCode string) models.School
	Create(school models.School) models.School
	DeleteSchool(id int) error
	EditSchool(id int, school models.School) error
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
	schoolRepo.DB.Raw("SELECT * FROM schools").Scan(&schools)
	return schools
}

func (schoolRepo *SchoolRepository) DeleteSchool(id int) error {
	var school models.School

	schoolRepo.DB.Unscoped().Delete(&school, id)

	return nil
}

func (schoolRepo *SchoolRepository) EditSchool(id int, school models.School) error {
	var sch models.School

	schoolRepo.DB.Model(&sch).Select("SchoolName", "Address", "Email", "Phone", "Website", "About",
		"StateGovStatus", "FederalGovernmentStatus", "CacRegNumber", "AssociationDetails", "SchoolCode",
		"LGACode", "LGACode", "Ownership", "LGA", "Logo", "SchoolTypes", "SchoolOperationTypes").Where("id", id).Updates(models.School{
		Model:                   gorm.Model{UpdatedAt: time.Now()},
		SchoolName:              school.SchoolName,
		Address:                 school.Address,
		Email:                   school.Email,
		Phone:                   school.Phone,
		Website:                 school.Website,
		About:                   school.About,
		StateGovStatus:          school.StateGovStatus,
		FederalGovernmentStatus: school.FederalGovernmentStatus,
		CacRegNumber:            school.CacRegNumber,
		AssociationDetails:      school.AssociationDetails,
		SchoolCode:              school.SchoolCode,
		LGACode:                 school.LGACode,
		Ownership:               school.Ownership,
		LGA:                     school.LGA,
		Logo:                    school.Logo,
		SchoolTypes:             school.SchoolTypes,
		SchoolOperationTypes:    school.SchoolOperationTypes,
	})

	return nil
}

func (schoolRepo *SchoolRepository) GetSchoolBySchoolCode(schoolCode string) models.School {
	var school models.School
	schoolRepo.DB.Unscoped().Table("schools").Where("school_code=?", schoolCode).Find(&school)
	return school
}
