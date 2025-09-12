package web

import (
	"GoNote/models"

	"github.com/gin-gonic/gin"
)

func (ws *WebServer) SessionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Читаем session_id из куки
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

			// Ставим куку на 7 дней
			c.SetCookie("session_id", sessionID, 7*24*3600, "/", "", false, true)
		} else {
			// Загружаем существующую сессию
			sess, err = ws.store.LoadSession(sessionID)
			if err != nil {
				c.AbortWithStatusJSON(500, gin.H{"error": "error loading session"})
				return
			}
		}

		// Сохраняем сессию в контекст Gin для использования в хендлерах
		c.Set("session", sess)
		c.Set("store", ws.store)
		c.Next()
	}
}
