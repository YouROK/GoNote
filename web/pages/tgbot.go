package pages

import (
	"GoNote/config"
	"GoNote/localize"
	"GoNote/tgbot"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ReportRequest struct {
	Reason string `json:"reason" binding:"required"`
	EMail  string `json:"email" binding:"required"`
	Text   string `json:"text" binding:"required"`
	Link   string `json:"link" binding:"required"`
}

func TGBotReport(c *gin.Context) {
	if config.Cfg.Features.DisableReportButton {
		c.JSON(http.StatusBadRequest, gin.H{"error": localize.T(c, "MsgErrDisableOnSite")})
		return
	}
	var req ReportRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("Error binding request report:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": localize.T(c, "MsgErrBadRequest")})
		return
	}

	message := fmt.Sprintf(
		localize.T(c, "TGBotMsgNewReport")+"\n\nReason: %s\nText: %s\nLink: %s\nEmail: %s",
		req.Reason,
		req.Text,
		req.Link,
		req.EMail,
	)

	tgbot.SendMessageAll(message)

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
