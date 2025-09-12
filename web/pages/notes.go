package pages

import (
	"bytes"
	"html/template"
	"math/rand"
	"net/http"
	"time"

	"GoNote/models"
	"GoNote/storage/fstorage"

	"github.com/gin-gonic/gin"
	"github.com/yuin/goldmark"
)

// HandleNote выдаёт страницу заметки
func HandleNote(c *gin.Context) {
	sess := c.MustGet("session").(*models.Session)
	store := c.MustGet("store").(*fstorage.FileStore)

	// Получаем noteID из URL
	noteID := c.Param("noteID")

	// Загружаем заметку
	note, content, err := store.GetNote(noteID)
	if err != nil {
		c.String(http.StatusNotFound, "Заметка не найдена")
		return
	}

	hasEdit := contains(sess.Notes, noteID)
	hasPass := note.Password != ""

	var buf bytes.Buffer
	if err := goldmark.Convert([]byte(content), &buf); err != nil {
		c.String(http.StatusInternalServerError, "Ошибка конвертации Markdown")
		return
	}
	htmlContent := template.HTML(buf.String())

	c.HTML(http.StatusOK, "note.go.html", gin.H{
		"note":    note,
		"content": htmlContent,
		"hasEdit": hasEdit,
		"hasPass": hasPass,
	})
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func CheckNotePassword(c *gin.Context) {
	time.Sleep(time.Millisecond * time.Duration(2000+rand.Intn(1000)))
	sess := c.MustGet("session").(*models.Session)
	store := c.MustGet("store").(*fstorage.FileStore)

	noteID := c.Param("noteID")

	var req struct {
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	note, _, err := store.GetNote(noteID)
	if err != nil || note == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "note not found"})
		return
	}

	if note.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "note is read only"})
		return
	}

	//TODO переделать на проверку хэша
	if req.Password == note.Password {
		exists := false
		for _, id := range sess.Notes {
			if id == noteID {
				exists = true
				break
			}
		}

		if !exists {
			sess.Notes = append(sess.Notes, noteID)
			store.SaveSession(sess)
		}
		c.JSON(http.StatusOK, gin.H{"ok": true})
		return
	}

	c.JSON(http.StatusUnauthorized, gin.H{"ok": false, "error": "invalid password"})
}

func EditNote(c *gin.Context) {
	sess := c.MustGet("session").(*models.Session)
	store := c.MustGet("store").(*fstorage.FileStore)

	noteID := c.Param("noteID")

	// Проверяем доступ в сессии
	hasAccess := false
	for _, id := range sess.Notes {
		if id == noteID {
			hasAccess = true
			break
		}
	}

	if !hasAccess {
		c.Redirect(http.StatusSeeOther, "/note/"+noteID)
		return
	}

	// Загружаем заметку
	note, content, err := store.GetNote(noteID)
	if err != nil || note == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "note not found"})
		return
	}

	// Если всё ок → рендерим страницу редактирования
	c.HTML(http.StatusOK, "pubnote.go.html", gin.H{
		"note":    note,
		"content": content,
	})
}
