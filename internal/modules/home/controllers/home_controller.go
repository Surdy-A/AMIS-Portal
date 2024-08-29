package controllers

import (
	"net/http"

	SchoolService "github.com/Surdy-A/amis_portal/internal/modules/school/services"
	"github.com/Surdy-A/amis_portal/pkg/html"
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

func (controller *Controller) Index(c *gin.Context) {
	html.Render(c, http.StatusOK, "modules/home/html/index", gin.H{
		"title":   "Home page",
		"flashes": sessions.Flash(c, "old"),
	})
}
