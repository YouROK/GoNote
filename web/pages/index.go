package pages

import (
	"GoNote/config"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// Главная страница
func IndexPage(c *gin.Context) {
	acceptLanguage := c.GetHeader("Accept-Language")
	isRussian := strings.HasPrefix(acceptLanguage, "ru") || strings.Contains(acceptLanguage, "ru;")

	c.HTML(http.StatusOK, "edit_note.go.html", gin.H{
		"note":                      nil,
		"content":                   "",
		"isRussian":                 isRussian,
		"disablePasswordPublishing": config.Cfg.Features.DisablePasswordPublishing,
	})
}
