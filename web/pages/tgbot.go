package pages

import (
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
	var req ReportRequest

	// Парсим JSON тело запроса
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("Error binding request report:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing reason or text"})
		return
	}

	// Формируем сообщение для админов
	message := fmt.Sprintf(
		"📢 New complaint received!\n\nReason: %s\nText: %s\nLink: %s\nEmail: %s",
		req.Reason,
		req.Text,
		req.Link,
		req.EMail,
	)

	// Отправляем всем админам
	tgbot.SendMessageAll(message)

	// Возвращаем успешный ответ
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
