package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Surdy-A/amis_portal/internal/modules/examination/models"
	ExaminationService "github.com/Surdy-A/amis_portal/internal/modules/examination/services"
	"github.com/Surdy-A/amis_portal/internal/modules/user/helpers"
	"github.com/Surdy-A/amis_portal/pkg/converters"
	"github.com/Surdy-A/amis_portal/pkg/errors"
	"github.com/Surdy-A/amis_portal/pkg/html"
	"github.com/Surdy-A/amis_portal/pkg/old"
	"github.com/Surdy-A/amis_portal/pkg/sessions"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	examinationService ExaminationService.ExaminationServiceInterface
}

func New() *Controller {
	return &Controller{
		examinationService: ExaminationService.New(),
	}
}

func (controller *Controller) GetExamination(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		html.Render(c, http.StatusInternalServerError, "templates/errors/html/500", gin.H{"title": "Server error", "message": "error converting the id"})
		return
	}

	examination, err := controller.examinationService.GetExamination(id)
	if err != nil {
		html.Render(c, http.StatusNotFound, "templates/errors/html/404", gin.H{"title": "Page not found", "message": err.Error()})
		return
	}

	html.Render(c, http.StatusOK, "modules/examination/html/examination_detail", gin.H{"title": "Show examination", "examination": examination})
}

func (controller *Controller) CreateExamination(c *gin.Context) {
	html.Render(c, http.StatusOK, "modules/examination/html/create_examination", gin.H{
		"title":  "Create Examination",
		"events": models.Events,
	})
}

func (controller *Controller) UpdateExamination(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		html.Render(c, http.StatusInternalServerError,
			"templates/errors/html/500",
			gin.H{"title": "Server error", "message": "error converting the id"})
		return
	}

	examination, err := controller.examinationService.GetExamination(id)
	if err != nil {
		html.Render(c, http.StatusNotFound, "templates/errors/html/404", gin.H{"title": "Page not found", "message": err.Error()})
		return
	}

	html.Render(c, http.StatusOK, "modules/examination/html/update_examination", gin.H{
		"title": "Update Examination",
		// "examination_type_lists":           models.ExaminationTypeList,
		// "examination_operation_type_lists": models.ExaminationOperationTypeLists,
		// "state_gov_status":            models.StateGovStatusList,
		// "federal_gov_status":          models.FederalGovStatusList,
		// "ownerships":                  models.Ownerships,
		// "LGANames":                    models.LGAName,
		// "examination":                 examination,
		"examination": examination,
	})
}

func (controller *Controller) Add(c *gin.Context) {
	user := helpers.Auth(c)

	// validate the request
	var examination models.Examination
	// This will infer what binder to use depending on the content-type header.
	if err := c.ShouldBind(&examination); err != nil {
		errors.Init()
		errors.SetFromErrors(err)
		sessions.Set(c, "errors", converters.MapToString(errors.Get()))

		old.Init()
		old.Set(c)
		sessions.Set(c, "old", converters.UrlValuesToString(old.Get()))

		c.Redirect(http.StatusFound, "/examinations/create")
		return
	}

	// Create the article
	examination, err := controller.examinationService.AddExamination(examination, user)

	// Check if there is any error on the examination creation
	if err != nil {
		c.Redirect(http.StatusFound, "/examinations/create")
		fmt.Println(err)
		return
	}

	c.Redirect(http.StatusFound, fmt.Sprintf("/examination/%d", examination.ID))
}

func (controller *Controller) GetExaminations(c *gin.Context) {
	examinations := controller.examinationService.GetExaminations()

	html.Render(c, http.StatusOK, "modules/examination/html/list_examinations", gin.H{"title": "Show examination", "examinations": examinations})
}

func (controller *Controller) DeleteExamination(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		html.Render(c, http.StatusInternalServerError,
			"templates/errors/html/500",
			gin.H{"title": "Server error", "message": "error converting the id"})
		return
	}

	err = controller.examinationService.DeleteExamination(id)

	if err != nil {
		c.Redirect(http.StatusFound, "/examinations")
		fmt.Println(err)
		return
	}

	c.Redirect(http.StatusFound, ("/"))
}

func (controller *Controller) EditExamination(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		html.Render(c, http.StatusInternalServerError,
			"templates/errors/html/500",
			gin.H{"title": "Server error", "message": "error converting the id"})
		return
	}

	// validate the request
	var examination models.Examination
	// This will infer what binder to use depending on the content-type header.
	if err := c.ShouldBind(&examination); err != nil {
		errors.Init()
		errors.SetFromErrors(err)
		sessions.Set(c, "errors", converters.MapToString(errors.Get()))

		old.Init()
		old.Set(c)
		sessions.Set(c, "old", converters.UrlValuesToString(old.Get()))

		c.Redirect(http.StatusFound, "/examinations/edit/{{id}}")
		return
	}
	err = controller.examinationService.EditExamination(id, examination)

	if err != nil {
		c.Redirect(http.StatusFound, "/examinations")
		fmt.Println(err)
		return
	}

	c.Redirect(http.StatusFound, fmt.Sprintf("/"))
}
