package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/Surdy-A/amis_portal/internal/modules/examination/models"
	ExaminationService "github.com/Surdy-A/amis_portal/internal/modules/examination/services"
	SchoolService "github.com/Surdy-A/amis_portal/internal/modules/school/services"
	"github.com/Surdy-A/amis_portal/internal/modules/user/helpers"
	"github.com/Surdy-A/amis_portal/pkg/converters"
	"github.com/Surdy-A/amis_portal/pkg/errors"
	"github.com/Surdy-A/amis_portal/pkg/html"
	"github.com/Surdy-A/amis_portal/pkg/old"
	"github.com/Surdy-A/amis_portal/pkg/sessions"
	"github.com/Surdy-A/amis_portal/pkg/toast"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	examinationService ExaminationService.ExaminationServiceInterface
	schoolService      SchoolService.SchoolServiceInterface
}

func New() *Controller {
	return &Controller{
		examinationService: ExaminationService.New(),
		schoolService:      SchoolService.New(),
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

func (controller *Controller) CreatePrimaryCatComp(c *gin.Context) {
	html.Render(c, http.StatusOK, "modules/examination/html/create_primary_cat_exam", gin.H{
		"title":   "Primary Category Examination",
		"events":  models.Events,
		"schools": controller.schoolService.GetSchools(),
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

func (controller *Controller) AddPrimaryCatComp(c *gin.Context) {
	user := helpers.Auth(c)

	// validate the request
	var examination models.PrimaryCompetition
	// This will infer what binder to use depending on the content-type header.
	if err := c.ShouldBind(&examination); err != nil {
		errors.Init()
		errors.SetFromErrors(err)
		sessions.Set(c, "errors", converters.MapToString(errors.Get()))

		old.Init()
		old.Set(c)
		sessions.Set(c, "old", converters.UrlValuesToString(old.Get()))

		c.Redirect(http.StatusFound, "/examination/create")
		return
	}

	// Create the article
	examination, err := controller.examinationService.AddPrimaryCatComp(examination, user)

	// Check if there is any error on the examination creation
	if err != nil {
		c.Redirect(http.StatusFound, "/examination/create")
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

	c.Redirect(http.StatusFound, "/")
}

func (controller *Controller) GetExamRegistration(c *gin.Context) {
	schools := controller.schoolService.GetSchools()
	html.Render(c, http.StatusOK, "modules/examination/html/create_exam_registration", gin.H{
		"title":      " Registration",
		"exam_lists": models.ExaminationList,
		"year":       time.Now().Year(),
		"schools":    schools,
	})
}

func (controller *Controller) GetCandidateInfo(c *gin.Context) {
	schoolCode := c.Param("id")
	fmt.Println(schoolCode)
	school, err := controller.schoolService.GetSchoolBySchoolCode(schoolCode)
	if err != nil {
		html.Render(c, http.StatusNotFound, "templates/errors/html/404", gin.H{"title": "Page not found", "message": err.Error()})
		return
	}

	examCode := c.Param("id")
	candidates, err := controller.examinationService.GetExaminationsBySchoolCode(examCode)
	if err != nil {
		html.Render(c, http.StatusNotFound, "templates/errors/html/404", gin.H{"title": "Page not found", "message": err.Error()})
		return
	}

	// noOfCandidates := len(candidates)
	// percentage := noOfCandidates/20*100
	html.Render(c, http.StatusOK, "modules/examination/html/candidate_info", gin.H{
		"title":            "Candidate Information",
		"school":           school,
		"candidates":       candidates,
		"no_of_candidates": len(candidates),
		//"percentage":       percentage,
	})
}

func (controller *Controller) CreateGradingExam(c *gin.Context) {
	user := helpers.Auth(c)
	// validate the request
	var gExam models.GradingExamination
	// This will infer what binder to use depending on the content-type header.
	if err := c.ShouldBind(&gExam); err != nil {
		errors.Init()
		errors.SetFromErrors(err)
		sessions.Set(c, "errors", converters.MapToString(errors.Get()))

		old.Init()
		old.Set(c)
		sessions.Set(c, "old", converters.UrlValuesToString(old.Get()))

		c.Redirect(http.StatusFound, "/examination/grading")
		return
	}
	_, err := controller.examinationService.CreateGradingExam(gExam, user)
	// Create the article
	// Check if there is any error on the examination creation
	if err != nil {
		c.Redirect(http.StatusFound, "/examination/error")
		return
	}

	toast.Success(c, "examination created successfully")
	c.Redirect(http.StatusFound, fmt.Sprintf("/examination/candidate/%d", int(gExam.ID)))
}

func (controller *Controller) AddCandidate(c *gin.Context) {
	schools := controller.schoolService.GetSchools()
	html.Render(c, http.StatusOK, "modules/examination/html/add_candidate", gin.H{
		"title":      "Add Candiadte",
		"exam_lists": models.ExaminationList,
		"lgas":       models.LGA,
		"year":       time.Now().Year(),
		"schools":    schools,
		"gender":     models.Gender,
	})
}

func (controller *Controller) CreateCandidate(c *gin.Context) {
	user := helpers.Auth(c)
	// validate the request
	var gExam models.StudentGradingExamInfo
	// This will infer what binder to use depending on the content-type header.
	if err := c.ShouldBind(&gExam); err != nil {
		errors.Init()
		errors.SetFromErrors(err)
		sessions.Set(c, "errors", converters.MapToString(errors.Get()))

		old.Init()
		old.Set(c)
		sessions.Set(c, "old", converters.UrlValuesToString(old.Get()))

		c.Redirect(http.StatusFound, "/examination/add-candidate")
		return
	}

	// Create the article
	_, err := controller.examinationService.CreateCandidate(gExam, user)
	// Check if there is any error on the examination creation
	if err != nil {
		c.Redirect(http.StatusFound, "/examination/candidates")
		fmt.Println(err)
		return
	}

	c.Redirect(http.StatusFound, "/examination/add-candidate")
}

func (controller *Controller) DisplayError(c *gin.Context) {
	html.Render(c, http.StatusOK, "modules/examination/html/error", gin.H{
		"title": "Error",
	})
}

func (controller *Controller) GetBlogPosts(c *gin.Context) {
	posts := controller.examinationService.GetBlogPosts()

	html.Render(c, http.StatusOK, "modules/examination/html/blog", gin.H{
		"title": "Blog",
		"posts": posts,
	})
}

func (controller *Controller) GetBlogPost(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	post, err := controller.examinationService.GetBlogPost(id)
	if err != nil {
		c.Redirect(http.StatusFound, "/examination/error")
		return
	}

	html.Render(c, http.StatusOK, "modules/examination/html/blog_detail", gin.H{
		"title": "Blog",
		"post":  post,
	})
}
