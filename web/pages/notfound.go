package pages

import "github.com/gin-gonic/gin"

func NotFound(c *gin.Context) {
	c.HTML(404, "notfound.go.html", nil)
}
