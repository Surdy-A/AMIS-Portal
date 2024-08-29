package routes

import (
	"github.com/Surdy-A/amis_portal/internal/middlewares"
	userCtrl "github.com/Surdy-A/amis_portal/internal/modules/user/controllers"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	userController := userCtrl.New()

	guestGroup := router.Group("/")
	guestGroup.Use(middlewares.IsGuest())
	{
		guestGroup.GET("/register", userController.Register)
		guestGroup.POST("/register", userController.HandleRegister)

		guestGroup.GET("/login", userController.Login)
		guestGroup.POST("/login", userController.HandleLogin)

		guestGroup.GET("/register-exam", userController.RegisterExam)
		guestGroup.POST("/register-exam", userController.HandleRegisterExam)
	}

	authGroup := router.Group("/")
	authGroup.Use(middlewares.IsAuth())
	{
		authGroup.POST("/logout", userController.HandleLogout)
	}
}
