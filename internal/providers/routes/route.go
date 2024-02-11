package routes

import (
	homeRoutes "github.com/Surdy-A/amis_portal/internal/modules/home/routes"
	schoolRoutes "github.com/Surdy-A/amis_portal/internal/modules/school/routes"
	examinationRoutes "github.com/Surdy-A/amis_portal/internal/modules/examination/routes"
	userRoutes "github.com/Surdy-A/amis_portal/internal/modules/user/routes"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	homeRoutes.Routes(router)
	userRoutes.Routes(router)
	schoolRoutes.Routes(router)
	examinationRoutes.Routes(router)
}
