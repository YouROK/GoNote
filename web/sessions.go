package web

import (
	"GoNote/models"
	"time"

	"github.com/gin-gonic/gin"
)

func (ws *WebServer) SessionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionID, err := c.Cookie("session_id")
		var sess *models.Session

		if err != nil || !ws.store.SessionExists(sessionID) {
			// Генерируем новую сессию
			sess, err = ws.store.CreateSession()
			if err != nil {
				c.AbortWithStatusJSON(500, gin.H{"error": "error creating session"})
				return
			}
			sessionID = sess.ID
		} else {
			// Загружаем существующую сессию
			sess, err = ws.store.LoadSession(sessionID)
			if err != nil {
				// Просрочена или битая → создаём новую
				sess, err = ws.store.CreateSession()
				if err != nil {
					c.AbortWithStatusJSON(500, gin.H{"error": "error creating session"})
					return
				}
				sessionID = sess.ID
			}
		}

		// Продлеваем время жизни
		sess.ExpiresAt = time.Now().Add(7 * 24 * time.Hour)
		_ = ws.store.SaveSession(sess)

		// Ставим/обновляем куку
		c.SetCookie("session_id", sessionID, 7*24*3600, "/", "", false, true)

		c.Set("session", sess)
		c.Set("store", ws.store)
		c.Next()
	}
}
