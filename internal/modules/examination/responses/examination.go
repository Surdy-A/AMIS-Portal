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
