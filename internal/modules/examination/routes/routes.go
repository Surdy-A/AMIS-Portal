package routes

import (
	"github.com/Surdy-A/amis_portal/internal/middlewares"
	examCtrl "github.com/Surdy-A/amis_portal/internal/modules/examination/controllers"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	examinationController := examCtrl.New()
	//router.GET("/schools", examinationController.GetSchools)
	//schoolGroup := router.Group("/school")
	//schoolGroup.GET("/:id", examinationController.GetSchool)

	authGroup := router.Group("/examination")
	authGroup.Use(middlewares.IsAuth())
	{
		authGroup.GET("/create", examinationController.CreateExamination)
		authGroup.POST("/add", examinationController.Add)
		authGroup.GET("/edit/:id", examinationController.UpdateExamination)
		authGroup.POST("/update/:id", examinationController.EditExamination)
		authGroup.GET("/delete/:id", examinationController.DeleteExamination)
	}
}
