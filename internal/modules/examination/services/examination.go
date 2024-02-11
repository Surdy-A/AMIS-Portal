package services

import (
	"errors"

	"github.com/Surdy-A/amis_portal/internal/modules/examination/models"
	ExaminationRepository "github.com/Surdy-A/amis_portal/internal/modules/examination/repositories"
	requests "github.com/Surdy-A/amis_portal/internal/modules/examination/requests"
	ExaminationResponse "github.com/Surdy-A/amis_portal/internal/modules/examination/responses"
	UserResponse "github.com/Surdy-A/amis_portal/internal/modules/user/responses"
)

type ExaminationServiceInterface interface {
	GetExaminations() []models.Examination
	GetExamination(id int) (models.Examination, error)
	StoreAsUser(request requests.StoreRequest, user UserResponse.User) (ExaminationResponse.Examination, error)
	AddExamination(sch models.Examination, user UserResponse.User) (models.Examination, error)
	DeleteExamination(id int) error
	EditExamination(id int, examination models.Examination) error
}

type ExaminationService struct {
	examinationRepository ExaminationRepository.ExaminationRepositoryInterface
}

func New() *ExaminationService {
	return &ExaminationService{
		examinationRepository: ExaminationRepository.New(),
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

func (examinationService *ExaminationService) StoreAsUser(examinationrequest requests.StoreRequest, user UserResponse.User) (ExaminationResponse.Examination, error) {
	var examination models.Examination
	var response ExaminationResponse.Examination

	examination.UserID = user.ID

	newExamination := examinationService.examinationRepository.Create(examination)

	if newExamination.ID == 0 {
		return response, errors.New("error in creating the examination")
	}

	return ExaminationResponse.ToExamination(newExamination), nil
}

func (examinationService *ExaminationService) AddExamination(sch models.Examination, user UserResponse.User) (models.Examination, error) {
	var examination models.Examination

	examination.Address = sch.Address
	// examination.Email = sch.Email
	// examination.Phone = sch.Phone
	// examination.Website = sch.Website
	// examination.About = sch.About
	// examination.StateGovStatus = sch.StateGovStatus
	// examination.FederalGovernmentStatus = sch.FederalGovernmentStatus
	// examination.CacRegNumber = sch.CacRegNumber
	// examination.AssociationDetails = sch.AssociationDetails
	// examination.LGACode = sch.LGACode
	// examination.Ownership = sch.Ownership
	// examination.LGA = sch.LGA
	// examination.Logo = sch.Logo
	examination.Paid = false

	examination.UserID = user.ID

	newExamination := examinationService.examinationRepository.Create(examination)

	if newExamination.ID == 0 {
		return models.Examination{}, errors.New("error in creating the examination")
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
