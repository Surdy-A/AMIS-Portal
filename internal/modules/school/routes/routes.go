package routes

import (
	"github.com/Surdy-A/amis_portal/internal/middlewares"
	schoolCtrl "github.com/Surdy-A/amis_portal/internal/modules/school/controllers"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	schoolController := schoolCtrl.New()
	router.GET("/schools", schoolController.GetSchools)
	schoolGroup := router.Group("/school")
	schoolGroup.GET("/:id", schoolController.GetSchool)

	authGroup := router.Group("/schools")
	authGroup.Use(middlewares.IsAuth())
	{
		authGroup.GET("/create", schoolController.CreateSchool)
		authGroup.POST("/add", schoolController.Add)
		authGroup.GET("/edit/:id", schoolController.UpdateSchool)
		authGroup.POST("/update/:id", schoolController.EditSchool)
		authGroup.GET("/delete/:id", schoolController.DeleteSchool)
	}
}
