package pages

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Главная страница
func HandleIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "pub_note.go.html", gin.H{
		"note":    nil,
		"content": "",
	})
}
