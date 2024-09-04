package services

import (
	"errors"

	"github.com/Surdy-A/amis_portal/internal/modules/examination/models"
	ExaminationRepository "github.com/Surdy-A/amis_portal/internal/modules/examination/repositories"
	requests "github.com/Surdy-A/amis_portal/internal/modules/examination/requests"
	ExaminationResponse "github.com/Surdy-A/amis_portal/internal/modules/examination/responses"
	SchoolRepository "github.com/Surdy-A/amis_portal/internal/modules/school/repositories"
	UserResponse "github.com/Surdy-A/amis_portal/internal/modules/user/responses"
	"gorm.io/gorm"
)

type ExaminationServiceInterface interface {
	GetExaminations() []models.Examination
	GetExamination(id int) (models.Examination, error)
	StoreAsUser(request requests.StoreRequest, user UserResponse.User) (ExaminationResponse.PrimaryCompetition, error)
	AddPrimaryCatComp(sch models.PrimaryCompetition, user UserResponse.User) (models.PrimaryCompetition, error)
	CreateGradingExam(sch models.GradingExamination, user UserResponse.User) (models.GradingExamination, error)
	CreateCandidate(sch models.StudentGradingExamInfo, user UserResponse.User) (models.StudentGradingExamInfo, error)
	DeleteExamination(id int) error
	EditExamination(id int, examination models.Examination) error
	GetExaminationsBySchoolCode(schoolCode string) ([]models.GradingExamination, error)
	GetExaminationBySchoolCode(schoolCode string) (models.GradingExamination, error)
	GetBlogPosts() []models.Blog
	GetBlogPost(id int) (models.Blog, error)
}

type ExaminationService struct {
	examinationRepository ExaminationRepository.ExaminationRepositoryInterface
	schoolRepository      SchoolRepository.SchoolRepositoryInterface
}

func New() *ExaminationService {
	return &ExaminationService{
		examinationRepository: ExaminationRepository.New(),
		schoolRepository:      SchoolRepository.New(),
	}
}

func (examinationService *ExaminationService) GetExaminations() []models.Examination {
	examinations := examinationService.examinationRepository.GetExaminations(400)
	for _, examination := range examinations {
		examinations = append(examinations, examination)
	}
	return examinations
}

func (examinationService *ExaminationService) GetExamination(id int) (models.Examination, error) {
	examination := examinationService.examinationRepository.GetExamination(id)

	if examination.ID == 0 {
		return models.Examination{}, errors.New("examination not found")
	}

	return examination, nil
}

func (examinationService *ExaminationService) StoreAsUser(examinationrequest requests.StoreRequest, user UserResponse.User) (ExaminationResponse.PrimaryCompetition, error) {
	var examination models.PrimaryCompetition
	var response ExaminationResponse.PrimaryCompetition

	examination.UserID = user.ID

	newExamination := examinationService.examinationRepository.AddPrimaryCatComp(examination)

	if newExamination.ID == 0 {
		return response, errors.New("error in creating the examination")
	}

	return ExaminationResponse.ToPrimaryCompetition(newExamination), nil
}

func (examinationService *ExaminationService) AddPrimaryCatComp(sch models.PrimaryCompetition, user UserResponse.User) (models.PrimaryCompetition, error) {
	examination := models.PrimaryCompetition{
		Quran:        "",
		Quiz_1:       "",
		Quiz_2:       "",
		Khutbah:      "",
		Exhibition_1: "",
		Exhibition_2: "",
		Examination: models.Examination{
			Model:              gorm.Model{},
			UserID:             user.ID,
			User:               sch.User,
			Address:            sch.Address,
			Zone:               sch.Zone,
			SchoolPhoneNumber:  sch.SchoolPhoneNumber,
			TeacherName:        sch.TeacherName,
			TeacherPhoneNumber: sch.TeacherName,
			HeadTeacherName:    sch.HeadTeacherName,
			LGA:                sch.LGA,
			Paid:               false,
			Category:           []string{},
		},
	}

	newExamination := examinationService.examinationRepository.AddPrimaryCatComp(examination)
	if newExamination.ID == 0 {
		return models.PrimaryCompetition{}, errors.New("error in creating the examination")
	}

	return newExamination, nil
}

func (examinationService *ExaminationService) DeleteExamination(id int) error {
	err := examinationService.examinationRepository.DeleteExamination(id)

	if err != nil {
		return errors.New("examination not found")
	}

	return nil
}

func (examinationService *ExaminationService) EditExamination(id int, examination models.Examination) error {
	err := examinationService.examinationRepository.EditExamination(id, examination)

	if err != nil {
		return errors.New("examination not found")
	}

	return nil
}

func (examinationService *ExaminationService) CreateGradingExam(sch models.GradingExamination, user UserResponse.User) (models.GradingExamination, error) {
	gExam := models.GradingExamination{
		UserID:     user.ID,
		ExamType:   sch.ExamType,
		SchoolCode: sch.SchoolCode,
		Surname:    sch.Surname,
		FirstName:  sch.FirstName,
		LastName:   sch.LastName,
		Gender:     sch.Gender,
		Age:        sch.Age,
	}

	ss := examinationService.schoolRepository.GetSchoolBySchoolCode(sch.SchoolCode)
	if gExam.SchoolCode != ss.SchoolCode {
		return models.GradingExamination{}, errors.New("invalid school code")
	}

	registeredExam, err := examinationService.GetExaminationsBySchoolCode(gExam.SchoolCode)
	if err != nil {
		return models.GradingExamination{}, err
	}

	var newExamination models.GradingExamination
	numberOfRegistredExam := len(registeredExam)
	if ss.Quota == 0 || numberOfRegistredExam >= ss.Quota {
		return models.GradingExamination{}, errors.New("no exam quota defined or exceded")
	} else {
		newExamination = examinationService.examinationRepository.CreateGradingExam(gExam)

		if newExamination.ID == 0 {
			return models.GradingExamination{}, errors.New("error in creating the examination")
		}

		return newExamination, nil
	}

}

func (examinationService *ExaminationService) CreateCandidate(sch models.StudentGradingExamInfo, user UserResponse.User) (models.StudentGradingExamInfo, error) {
	//var gExam models.StudentGradingExamInfo
	//gExam.Quota = 28
	sch.Quota = 28

	//gExam.SchoolCode = sch.SchoolCode
	ss := examinationService.schoolRepository.GetSchoolBySchoolCode(sch.SchoolCode)
	if sch.SchoolCode != ss.SchoolCode {
		return models.StudentGradingExamInfo{}, errors.New("invalid school code")
	}

	// gExam.UserID = user.ID
	// gExam.ExamType = sch.ExamType

	newExamination := examinationService.examinationRepository.CreateCandidate(sch)

	if newExamination.ID == 0 {
		return models.StudentGradingExamInfo{}, errors.New("error in creating the examination")
	}
	return newExamination, nil
}

func (examinationService *ExaminationService) GetExaminationsBySchoolCode(schoolCode string) ([]models.GradingExamination, error) {
	//var examination models.GradingExamination
	examinations, err := examinationService.examinationRepository.GetExaminationsBySchoolCode(schoolCode)
	if err != nil {
		return []models.GradingExamination{}, errors.New("examination not found")
	}

	//examinations = append(examinations, examination)

	return examinations, nil
}

func (examinationService *ExaminationService) GetExaminationBySchoolCode(schoolCode string) (models.GradingExamination, error) {
	//var examination models.GradingExamination
	examination, err := examinationService.examinationRepository.GetExaminationBySchoolCode(schoolCode)
	if err != nil {
		return models.GradingExamination{}, errors.New("examination not found")
	}

	return examination, nil
}

func (examinationService *ExaminationService) GetBlogPosts() []models.Blog {
	blogPosts := examinationService.examinationRepository.GetBlogPosts()
	return blogPosts
}

func (examinationService *ExaminationService) GetBlogPost(id int) (models.Blog, error) {
	post := examinationService.examinationRepository.GetBlogPost(id)

	if post.ID == 0 {
		return models.Blog{}, errors.New("post not found")
	}

	return post, nil
}
