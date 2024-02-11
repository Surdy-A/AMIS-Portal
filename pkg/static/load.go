package static

import (
	"github.com/gin-gonic/gin"
)

func LoadStatic(router *gin.Engine) {
	router.Static("/assets", "./assets")
	router.Static("/schools/assets", "./assets")
	router.Static("/school/assets", "./assets")
	router.Static("/schools/edit/assets", "./assets")
	//examination route
	router.Static("/examinations/assets", "./assets")
	router.Static("/examination/assets", "./assets")
	router.Static("/examination/edit/assets", "./assets")
}
