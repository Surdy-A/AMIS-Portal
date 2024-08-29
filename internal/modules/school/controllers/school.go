package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Surdy-A/amis_portal/internal/modules/school/models"
	SchoolService "github.com/Surdy-A/amis_portal/internal/modules/school/services"
	"github.com/Surdy-A/amis_portal/internal/modules/user/helpers"
	"github.com/Surdy-A/amis_portal/pkg/converters"
	"github.com/Surdy-A/amis_portal/pkg/errors"
	"github.com/Surdy-A/amis_portal/pkg/html"
	"github.com/Surdy-A/amis_portal/pkg/old"
	"github.com/Surdy-A/amis_portal/pkg/sessions"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	schoolService SchoolService.SchoolServiceInterface
}

func New() *Controller {
	return &Controller{
		schoolService: SchoolService.New(),
	}
}

func (controller *Controller) GetSchool(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		html.Render(c, http.StatusInternalServerError, "templates/errors/html/500", gin.H{"title": "Server error", "message": "error converting the id"})
		return
	}

	school, err := controller.schoolService.GetSchool(id)
	if err != nil {
		html.Render(c, http.StatusNotFound, "templates/errors/html/404", gin.H{"title": "Page not found", "message": err.Error()})
		return
	}

	html.Render(c, http.StatusOK, "modules/school/html/school_detail", gin.H{"title": "Show school", "school": school})
}

func (controller *Controller) CreateSchool(c *gin.Context) {
	html.Render(c, http.StatusOK, "modules/school/html/create_school", gin.H{
		"title":                       "Create School",
		"school_type_lists":           models.SchoolTypeList,
		"school_operation_type_lists": models.SchoolOperationTypeLists,
		"state_gov_status":            models.StateGovStatusList,
		"federal_gov_status":          models.FederalGovStatusList,
		"ownerships":                  models.Ownerships,
		"LGANames":                    models.LGAName,
	})
}

func (controller *Controller) UpdateSchool(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		html.Render(c, http.StatusInternalServerError,
			"templates/errors/html/500",
			gin.H{"title": "Server error", "message": "error converting the id"})
		return
	}

	school, err := controller.schoolService.GetSchool(id)
	if err != nil {
		html.Render(c, http.StatusNotFound, "templates/errors/html/404", gin.H{"title": "Page not found", "message": err.Error()})
		return
	}

	html.Render(c, http.StatusOK, "modules/school/html/update_school", gin.H{
		"title":                       "Update School",
		"school_type_lists":           models.SchoolTypeList,
		"school_operation_type_lists": models.SchoolOperationTypeLists,
		"state_gov_status":            models.StateGovStatusList,
		"federal_gov_status":          models.FederalGovStatusList,
		"ownerships":                  models.Ownerships,
		"LGANames":                    models.LGAName,
		"school":                      school,
	})
}

func (controller *Controller) Add(c *gin.Context) {
	user := helpers.Auth(c)

	// validate the request
	var sch models.School
	// This will infer what binder to use depending on the content-type header.
	if err := c.ShouldBind(&sch); err != nil {
		errors.Init()
		errors.SetFromErrors(err)
		sessions.Set(c, "errors", converters.MapToString(errors.Get()))

		old.Init()
		old.Set(c)
		sessions.Set(c, "old", converters.UrlValuesToString(old.Get()))

		c.Redirect(http.StatusFound, "/schools/create")
		return
	}

	// Create the article
	school, err := controller.schoolService.AddSchool(sch, user)

	// Check if there is any error on the school creation
	if err != nil {
		c.Redirect(http.StatusFound, "/schools/create")
		fmt.Println(err)
		return
	}

	c.Redirect(http.StatusFound, fmt.Sprintf("/school/%d", school.ID))
}

func (controller *Controller) GetSchools(c *gin.Context) {
	schools := controller.schoolService.GetSchools()

	html.Render(c, http.StatusOK, "modules/school/html/list_schools", gin.H{"title": "Show school", "schools": schools})
}

func (controller *Controller) DeleteSchool(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		html.Render(c, http.StatusInternalServerError,
			"templates/errors/html/500",
			gin.H{"title": "Server error", "message": "error converting the id"})
		return
	}

	err = controller.schoolService.DeleteSchool(id)

	if err != nil {
		c.Redirect(http.StatusFound, "/schools")
		fmt.Println(err)
		return
	}

	c.Redirect(http.StatusFound, ("/"))
}

func (controller *Controller) EditSchool(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		html.Render(c, http.StatusInternalServerError,
			"templates/errors/html/500",
			gin.H{"title": "Server error", "message": "error converting the id"})
		return
	}

	// validate the request
	var sch models.School
	// This will infer what binder to use depending on the content-type header.
	if err := c.ShouldBind(&sch); err != nil {
		errors.Init()
		errors.SetFromErrors(err)
		sessions.Set(c, "errors", converters.MapToString(errors.Get()))

		old.Init()
		old.Set(c)
		sessions.Set(c, "old", converters.UrlValuesToString(old.Get()))

		c.Redirect(http.StatusFound, "/schools/edit/{{id}}")
		return
	}
	err = controller.schoolService.EditSchool(id, sch)

	if err != nil {
		c.Redirect(http.StatusFound, "/schools")
		fmt.Println(err)
		return
	}

	c.Redirect(http.StatusFound, fmt.Sprintf("/"))
}

func (controller *Controller) GetSchoolBySchoolCode(c *gin.Context) {
	sc := c.Param("school_code")
	school, err := controller.schoolService.GetSchoolBySchoolCode(sc)
	if err != nil {
		html.Render(c, http.StatusNotFound, "templates/errors/html/404", gin.H{"title": "Page not found", "message": err.Error()})
		return
	}

	html.Render(c, http.StatusOK, "modules/school/html/school_detail", gin.H{"title": "Show school", "school": school})
}
