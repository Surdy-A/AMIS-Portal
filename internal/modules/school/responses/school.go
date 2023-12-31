package responses

import (
	"github.com/Surdy-A/amis_portal/internal/modules/school/models"
	UserResponse "github.com/Surdy-A/amis_portal/internal/modules/user/responses"
)

type School struct {
	models.School
	User UserResponse.User
}

type Schools struct {
	Data []School
}

func ToSchool(school models.School) School {
	return School{}
}

func ToSchools(schools []models.School) Schools {
	var response Schools

	for _, school := range schools {
		response.Data = append(response.Data, ToSchool(school))
	}

	return response
}
