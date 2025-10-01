package pages

import (
	"GoNote/config"
	"GoNote/localize"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Главная страница
func IndexPage(c *gin.Context) {
	c.HTML(http.StatusOK, "edit_note.go.html", localize.AddMessages(c, "edit_note.go.html", gin.H{
		"note":                      nil,
		"content":                   "",
		"menu":                      "",
		"disablePasswordPublishing": config.Cfg.Features.DisablePasswordPublishing,
	}))
}
