package pages

import (
	"GoNote/localize"

	"github.com/gin-gonic/gin"
)

func NotFound(c *gin.Context) {
	c.HTML(404, "notfound.go.html", localize.AddMessages(c, "notfound.go.html"))
}
