package repositories

import (
	"time"

	"github.com/Surdy-A/amis_portal/internal/modules/examination/models"
	"github.com/Surdy-A/amis_portal/pkg/database"
	"gorm.io/gorm"
)

type ExaminationRepositoryInterface interface {
	GetExaminations(limit int) []models.Examination
	GetExamination(id int) models.Examination
	Create(examination models.Examination) models.Examination
	DeleteExamination(id int) error
	EditExamination(id int, examination models.Examination) error
}

type ExaminationRepository struct {
	DB *gorm.DB
}

func New() *ExaminationRepository {
	return &ExaminationRepository{
		DB: database.Connection(),
	}
}

func (examinationRepo *ExaminationRepository) Create(examination models.Examination) models.Examination {
	var newExamination models.Examination

	examinationRepo.DB.Create(&examination).Scan(&newExamination)

	return newExamination
}

func (examinationRepo *ExaminationRepository) GetExamination(id int) models.Examination {
	var examination models.Examination

	examinationRepo.DB.First(&examination, id)

	return examination
}

func (examinationRepo *ExaminationRepository) GetExaminations(limit int) []models.Examination {
	var examinations []models.Examination

	examinationRepo.DB.Limit(limit).Joins("User").Find(&examinations)

	//examinations = append(examinations, sch)
	return examinations
}

func (examinationRepo *ExaminationRepository) DeleteExamination(id int) error {
	var examination models.Examination

	examinationRepo.DB.Unscoped().Delete(&examination, id)

	return nil
}

func (examinationRepo *ExaminationRepository) EditExamination(id int, examination models.Examination) error {
	var sch models.Examination

	examinationRepo.DB.Model(&sch).Select("ExaminationName").Where("id", id).Updates(models.Examination{
		Model: gorm.Model{UpdatedAt: time.Now()},
	})

	return nil
}
