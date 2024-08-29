package routes

import (
	homeCtrl "github.com/Surdy-A/amis_portal/internal/modules/home/controllers"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	homeController := homeCtrl.New()
	router.GET("/", homeController.Index)
//	router.GET("/blog", homeController.GetBlogPost)
	//	guestGroup.POST("/blog", userController.AddBlogPost)

}
