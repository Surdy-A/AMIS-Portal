package html

import (
	"github.com/Surdy-A/amis_portal/internal/providers/view"
	"github.com/gin-gonic/gin"
)

func Render(c *gin.Context, code int, name string, data gin.H) {
	data = view.WithGlobalData(c, data)

	c.HTML(code, name, data)
}
