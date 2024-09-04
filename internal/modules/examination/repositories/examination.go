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
	AddPrimaryCatComp(examination models.PrimaryCompetition) models.PrimaryCompetition
	CreateGradingExam(examination models.GradingExamination) models.GradingExamination
	CreateCandidate(examination models.StudentGradingExamInfo) models.StudentGradingExamInfo
	DeleteExamination(id int) error
	EditExamination(id int, examination models.Examination) error
	GetExaminationsBySchoolCode(schoolCode string) ([]models.GradingExamination, error)
	GetExaminationBySchoolCode(schoolCode string) (models.GradingExamination, error)
	GetBlogPosts() []models.Blog
	GetBlogPost(id int) models.Blog
}

type ExaminationRepository struct {
	DB *gorm.DB
}

func New() *ExaminationRepository {
	return &ExaminationRepository{
		DB: database.Connection(),
	}
}

func (examinationRepo *ExaminationRepository) AddPrimaryCatComp(examination models.PrimaryCompetition) models.PrimaryCompetition {
	var newExamination models.PrimaryCompetition

	examinationRepo.DB.Table("primary-competiton").Create(&examination).Scan(&newExamination)

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

func (examinationRepo *ExaminationRepository) CreateGradingExam(examination models.GradingExamination) models.GradingExamination {
	var newExamination models.GradingExamination

	examinationRepo.DB.Create(&examination).Scan(&newExamination)

	return newExamination
}

func (examinationRepo *ExaminationRepository) CreateCandidate(examination models.StudentGradingExamInfo) models.StudentGradingExamInfo {
	//var newExamination models.StudentGradingExamInfo
	newExamination := models.StudentGradingExamInfo{
		GradingExamination: models.GradingExamination{},
		Quota:              28,
		LastName:           examination.LastName,
		FirstName:          examination.FirstName,
		StateOfOrigin:      examination.StateOfOrigin,
		LGA:                examination.LGA,
		Gender:             examination.Gender,
		Added:              time.Now(),
	}
	examinationRepo.DB.Table("student_grading_exam_infos").Where("school_code=? AND id=?", examination.SchoolCode, examination.ID).Updates(&examination)

	return newExamination
}

func (examinationRepo *ExaminationRepository) GetExaminationsBySchoolCode(schoolCode string) ([]models.GradingExamination, error) {
	var examinations []models.GradingExamination
	examinationRepo.DB.Table("grading_examinations").Where("school_code=?", schoolCode).Find(&examinations)

	return examinations, nil
}

func (examinationRepo *ExaminationRepository) GetExaminationBySchoolCode(schoolCode string) (models.GradingExamination, error) {
	var examination models.GradingExamination
	examinationRepo.DB.Table("grading_examinations").Where("school_code=?", schoolCode).Find(&examination)
	return examination, nil
}

func (examinationRepo *ExaminationRepository) GetBlogPosts() []models.Blog {
	var blogPosts []models.Blog
	examinationRepo.DB.Raw("SELECT * FROM blogs").Scan(&blogPosts)

	return blogPosts
}

func (examinationRepo *ExaminationRepository) GetBlogPost(id int) models.Blog {
	var blogPost models.Blog
	examinationRepo.DB.First(&blogPost, id)

	return blogPost
}
