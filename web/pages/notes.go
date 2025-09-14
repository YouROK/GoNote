package pages

import (
	"GoNote/utils"
	"html/template"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"GoNote/models"
	"GoNote/storage/fstorage"

	"github.com/gin-gonic/gin"
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

	c.HTML(http.StatusOK, "view_note.go.html", gin.H{
		"note":    note,
		"content": template.HTML(content),
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
		time.Sleep(time.Millisecond * time.Duration(2000+rand.Intn(1000)))
		c.JSON(http.StatusBadRequest, gin.H{"error": "note is read only"})
		return
	}

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

	time.Sleep(time.Millisecond * time.Duration(2000+rand.Intn(1000)))
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

	acceptLanguage := c.GetHeader("Accept-Language")
	isRussian := strings.HasPrefix(acceptLanguage, "ru") || strings.Contains(acceptLanguage, "ru;")

	c.HTML(http.StatusOK, "pub_note.go.html", gin.H{
		"note":      note,
		"content":   template.HTML(content),
		"isRussian": isRussian,
	})
}

func PublishNote(c *gin.Context) {
	sess := c.MustGet("session").(*models.Session)
	store := c.MustGet("store").(*fstorage.FileStore)

	var note *models.Note
	var err error

	// Парсим данные из запроса
	var req struct {
		Title    string `json:"title"`
		Author   string `json:"author"`
		Content  string `json:"content"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	noteID := c.Param("noteID")
	isNewNote := noteID == "new"

	if !isNewNote {
		// Обновляем заметку

		// Смотрим, может ли пользователь редактировать
		hasAccess := false
		for _, id := range sess.Notes {
			if id == noteID {
				hasAccess = true
				break
			}
		}

		// Проверяем пароль
		note, _, err = store.GetNote(noteID)
		if err != nil || note == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "note not found"})
			return
		}
		if !hasAccess && note.Password != "" && req.Password != note.Password {
			// Пользователь не может редактировать и в заметке задан пароль и не совпадает
			time.Sleep(time.Millisecond * time.Duration(4000+rand.Intn(1000)))
			// Хреновая защита от подбора, нужно запросы от ip блокировать на время
			c.JSON(http.StatusForbidden, gin.H{"error": "wrong password"})
			return
		}

		if !hasAccess && (note.Password == "" || req.Password != note.Password) {
			c.JSON(http.StatusForbidden, gin.H{"error": "no access"})
			return
		}
	}

	// Новая заметка
	if note == nil {
		// Подбираем уникальный id
		id := ""
		for {
			id = utils.RandStr(8)
			if n, _, _ := store.GetNote(id); n == nil {
				break
			}
		}

		note = &models.Note{
			ID:        id,
			CreatedAt: time.Now(),
		}
	}

	// Обновляем поля
	note.Title = req.Title
	note.Author = req.Author
	if req.Password != "" {
		note.Password = req.Password
	}
	note.UpdatedAt = time.Now()

	// Сохраняем заметку
	err = store.UpdateNote(note, req.Content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save note"})
		return
	}

	// Запоминаем в сессии заметку
	exists := false
	for _, id := range sess.Notes {
		if id == note.ID {
			exists = true
			break
		}
	}

	if !exists {
		sess.Notes = append(sess.Notes, note.ID)
		store.SaveSession(sess)
	}

	c.JSON(http.StatusOK, gin.H{"ok": true, "noteID": note.ID})
}
