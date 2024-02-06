package services

import (
	"errors"

	"github.com/Surdy-A/amis_portal/internal/modules/school/models"
	SchoolRepository "github.com/Surdy-A/amis_portal/internal/modules/school/repositories"
	requests "github.com/Surdy-A/amis_portal/internal/modules/school/requests"
	SchoolResponse "github.com/Surdy-A/amis_portal/internal/modules/school/responses"
	UserResponse "github.com/Surdy-A/amis_portal/internal/modules/user/responses"
)

type SchoolServiceInterface interface {
	GetSchools() []models.School
	GetSchool(id int) (models.School, error)
	StoreAsUser(request requests.StoreRequest, user UserResponse.User) (SchoolResponse.School, error)
	AddSchool(sch models.School, user UserResponse.User) (models.School, error)
	DeleteSchool(id int) error
	EditSchool(id int, school models.School) error
}

type SchoolService struct {
	schoolRepository SchoolRepository.SchoolRepositoryInterface
}

func New() *SchoolService {
	return &SchoolService{
		schoolRepository: SchoolRepository.New(),
	}
}

func (schoolService *SchoolService) GetSchools() []models.School {
	//var schools []models.School

	schools := schoolService.schoolRepository.GetSchools(400)
	for _, school := range schools {
		schools = append(schools, school)
	}
	return schools
}

// func (schoolService *SchoolService) GetStoriesArticles() ArticleResponse.Articles {
// 	articles := articleService.articleRepository.List(6)

// 	return ArticleResponse.ToArticles(articles)
// }

func (schoolService *SchoolService) GetSchool(id int) (models.School, error) {
	school := schoolService.schoolRepository.GetSchool(id)

	if school.ID == 0 {
		return models.School{}, errors.New("school not found")
	}

	return school, nil
}

func (schoolService *SchoolService) StoreAsUser(schoolrequest requests.StoreRequest, user UserResponse.User) (SchoolResponse.School, error) {
	var school models.School
	var response SchoolResponse.School

	//school.SchoolName = schoolrequest.SchoolName
	school.UserID = user.ID

	newSchool := schoolService.schoolRepository.Create(school)

	if newSchool.ID == 0 {
		return response, errors.New("error in creating the school")
	}

	return SchoolResponse.ToSchool(newSchool), nil
}

func (schoolService *SchoolService) AddSchool(sch models.School, user UserResponse.User) (models.School, error) {
	var school models.School

	school.SchoolName = sch.SchoolName
	school.Address = sch.Address
	school.Email = sch.Email
	school.Phone = sch.Phone
	school.Website = sch.Website
	school.About = sch.About
	school.StateGovStatus = sch.StateGovStatus
	school.FederalGovernmentStatus = sch.FederalGovernmentStatus
	school.CacRegNumber = sch.CacRegNumber
	school.AssociationDetails = sch.AssociationDetails
	school.SchoolCode = sch.SchoolCode
	school.LGACode = sch.LGACode
	school.Ownership = sch.Ownership
	school.LGA = sch.LGA
	school.Logo = sch.Logo
	school.Paid = false
	school.SchoolTypes = sch.SchoolTypes
	school.SchoolOperationTypes = sch.SchoolOperationTypes

	school.UserID = user.ID

	newSchool := schoolService.schoolRepository.Create(school)

	if newSchool.ID == 0 {
		return models.School{}, errors.New("error in creating the school")
	}
	return newSchool, nil
}

func (schoolService *SchoolService) DeleteSchool(id int) error {
	err := schoolService.schoolRepository.DeleteSchool(id)

	if err != nil {
		return errors.New("school not found")
	}

	return nil
}

func (schoolService *SchoolService) EditSchool(id int, school models.School) error {
	err := schoolService.schoolRepository.EditSchool(id, school)

	if err != nil {
		return errors.New("school not found")
	}

	return nil
}
