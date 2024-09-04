package responses

import (
	"github.com/Surdy-A/amis_portal/internal/modules/examination/models"
	UserResponse "github.com/Surdy-A/amis_portal/internal/modules/user/responses"
)

type Examination struct {
	models.Examination
	User UserResponse.User
}

type Examinations struct {
	Data []Examination
}

type PrimaryCompetition struct {
	models.PrimaryCompetition
	User UserResponse.User
}

type PrimaryCompetitions struct {
	Data []PrimaryCompetition
}

func ToExamination(school models.Examination) Examination {
	return Examination{}
}

func ToExaminations(schools []models.Examination) Examinations {
	var response Examinations

	for _, school := range schools {
		response.Data = append(response.Data, ToExamination(school))
	}

	return response
}

func ToPrimaryCompetition(school models.PrimaryCompetition) PrimaryCompetition {
	return PrimaryCompetition{}
}

func ToPrimaryCompetitions(primaryCompetitions []models.PrimaryCompetition) PrimaryCompetitions {
	var response PrimaryCompetitions

	for _, primaryCompetition := range primaryCompetitions {
		response.Data = append(response.Data, ToPrimaryCompetition(primaryCompetition))
	}

	return response
}
