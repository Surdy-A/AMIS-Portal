package middlewares

import (
	UserRepository "github.com/Surdy-A/amis_portal/internal/modules/user/repositories"
	"github.com/Surdy-A/amis_portal/pkg/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func IsAuth() gin.HandlerFunc {
	var userRepo = UserRepository.New()

	return func(c *gin.Context) {
		authID := sessions.Get(c, "auth")
		userID, _ := strconv.Atoi(authID)

		user := userRepo.FindByID(userID)

		if user.ID == 0 {
			c.Redirect(http.StatusFound, "/login")
			return
		}
		// before request

		c.Next()
	}
}
