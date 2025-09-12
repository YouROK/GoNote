package pages

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// Главная страница
func HandleIndex(c *gin.Context) {
	acceptLanguage := c.GetHeader("Accept-Language")
	isRussian := strings.HasPrefix(acceptLanguage, "ru") || strings.Contains(acceptLanguage, "ru;")

	c.HTML(http.StatusOK, "pub_note.go.html", gin.H{
		"note":      nil,
		"content":   "",
		"isRussian": isRussian,
	})
}
