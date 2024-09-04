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
	router.GET("blog", examinationController.GetBlogPosts)
	router.GET("blog/:id", examinationController.GetBlogPost)
	authGroup := router.Group("/examination")
	authGroup.Use(middlewares.IsAuth())
	{
		authGroup.GET("/create", examinationController.CreatePrimaryCatComp)
		authGroup.POST("/add", examinationController.AddPrimaryCatComp)
		authGroup.GET("/edit/:id", examinationController.UpdateExamination)
		authGroup.PUT("/update/:id", examinationController.EditExamination)
		authGroup.GET("/delete/:id", examinationController.DeleteExamination)
		authGroup.GET("/grading", examinationController.GetExamRegistration)    //for get templatec
		authGroup.POST("/grading-add", examinationController.CreateGradingExam) //for post method
		authGroup.GET("/candidates/:id", examinationController.GetCandidateInfo)
		authGroup.GET("/add-candidate", examinationController.AddCandidate)        //for add candidate get templatec
		authGroup.POST("/create-candidate", examinationController.CreateCandidate) //for  post method
		authGroup.GET("/error", examinationController.DisplayError)

		//	authGroup.POST("/add", examinationController.CandidateInfo)
		//authGroup.POST("/", examinationController.CreateExamRegistration) // to register exam

	}
}
