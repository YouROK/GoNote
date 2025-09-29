package pages

import (
	"GoNote/config"
	"GoNote/storage"
	"GoNote/tgbot"
	"GoNote/utils"
	"GoNote/web/sanitizer"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"GoNote/models"

	"github.com/gin-gonic/gin"
	"github.com/rainycape/unidecode"

	"github.com/microcosm-cc/bluemonday"
)

func incrementCounter(c *gin.Context, noteID string, store storage.Store) int {
	h := md5.Sum([]byte(noteID))
	cookieName := "note_" + hex.EncodeToString(h[:]) + "_viewed"

	// Проверяем наличие cookie
	_, err := c.Cookie(cookieName)

	if err != nil {
		// cookie нет → увеличиваем счётчик
		count, err := store.IncrementCounterViews(noteID)
		if err != nil {
			log.Println("Error increment counter:", err)
			return 1
		}

		// ставим cookie на 1 час
		c.SetCookie(cookieName, "1", config.Cfg.Counter.TTLSeconds, "/", "", false, true)

		return count.Count
	}

	// cookie есть → просто возвращаем текущее значение
	count, err := store.GetCounterViews(noteID)
	if err != nil {
		log.Println("Error get counter:", err)
		return 1
	}

	return count.Count
}

// NotePage выдаёт страницу заметки
func NotePage(c *gin.Context) {
	sess := c.MustGet("session").(*models.Session)
	store := c.MustGet("store").(storage.Store)

	// Получаем noteID из URL
	noteID := c.Param("noteID")

	// Загружаем заметку
	note, content, menu, err := store.GetNote(noteID)
	if err != nil {
		c.String(http.StatusNotFound, "Заметка не найдена")
		return
	}

	hasEdit := contains(sess.Notes, noteID)
	hasPass := note.Password != ""

	counter := incrementCounter(c, noteID, store)

	c.HTML(http.StatusOK, "view_note.go.html", gin.H{
		"note":    note,
		"content": template.HTML(content),
		"menu":    template.HTML(menu),
		"hasEdit": hasEdit,
		"hasPass": hasPass,
		"counter": counter,
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
	store := c.MustGet("store").(storage.Store)

	noteID := c.Param("noteID")

	var req struct {
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	note, _, _, err := store.GetNote(noteID)
	if err != nil || note == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "note not found"})
		return
	}

	if note.Password == "" {
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

	c.JSON(http.StatusUnauthorized, gin.H{"ok": false, "error": "invalid password"})
}

func EditNotePage(c *gin.Context) {
	sess := c.MustGet("session").(*models.Session)
	store := c.MustGet("store").(storage.Store)

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
	note, content, menu, err := store.GetNote(noteID)
	if err != nil || note == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "note not found"})
		return
	}

	acceptLanguage := c.GetHeader("Accept-Language")
	isRussian := strings.HasPrefix(acceptLanguage, "ru") || strings.Contains(acceptLanguage, "ru;")

	c.HTML(http.StatusOK, "edit_note.go.html", gin.H{
		"note":      note,
		"content":   template.HTML(content),
		"menu":      template.HTML(menu),
		"isRussian": isRussian,
	})
}

type reqAddUpdNote struct {
	Title    string `json:"title"`
	Author   string `json:"author"`
	Content  string `json:"content"`
	Menu     string `json:"menu"`
	Password string `json:"password"`
}

func checkAddUpdNote(c *gin.Context) (*reqAddUpdNote, bool) {
	var req *reqAddUpdNote
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return nil, false
	}

	if len(req.Title) < 3 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title is too small"})
		return nil, false
	}

	// ================================
	// Сначала чистим inline style в span (KaTeX)
	// ================================
	req.Content = sanitizer.SanitizeStyleAttrs(req.Content, "span")

	// ================================
	// Политика белого списка для контента
	// ================================
	policy := bluemonday.NewPolicy()

	// Текстовые теги
	policy.AllowElements(
		"p", "h1", "h2", "h3", "h4",
		"strong", "b", "em", "i",
		"ul", "ol", "li",
		"hr", "br", "a", "img", "video",
	)
	policy.AllowAttrs("href").OnElements("a")
	policy.RequireParseableURLs(true)
	policy.AllowURLSchemes("http", "https")
	policy.AllowAttrs("src", "alt", "title").OnElements("img")
	policy.AllowAttrs("controls", "autoplay", "loop").OnElements("video")
	policy.AllowDataURIImages()

	// KaTeX / MathML
	policy.AllowElements("span")
	policy.AllowElements("math", "mrow", "mi", "mo", "msup", "mn", "semantics", "annotation")
	policy.AllowAttrs("class", "data-value", "contenteditable", "aria-hidden", "style").OnElements("span")
	policy.AllowAttrs("xmlns", "encoding").OnElements("math", "annotation")

	req.Content = policy.Sanitize(req.Content)

	// ================================
	// Политика для меню (обычный whitelist)
	// ================================
	policyMenu := bluemonday.UGCPolicy()
	policyMenu.AllowElements(
		"h2", "h3", "h4",
		"p", "a", "hr",
		"strong", "b",
		"em", "i",
	)
	policyMenu.AllowAttrs("href").OnElements("a")
	policyMenu.RequireParseableURLs(true)
	policyMenu.AllowURLSchemes("http", "https")
	policyMenu.AllowAttrs("class").Matching(
		regexp.MustCompile(`^ql-[a-z0-9\-]+$`),
	).OnElements("p", "h2", "h3", "h4", "ol", "ul", "li")
	req.Menu = policyMenu.Sanitize(req.Menu)

	// ================================
	// Ограничения по размеру
	// ================================
	if len(req.Content) > 1000000 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Content exceeds maximum size"})
		return nil, false
	}

	if len(req.Menu) > 10000 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Menu content exceeds maximum size"})
		return nil, false
	}

	if len(req.Content) == 0 || isEmptyContent(req.Content) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Content cannot be empty"})
		return nil, false
	}

	return req, true
}

func NewNote(c *gin.Context) {
	sess := c.MustGet("session").(*models.Session)
	store := c.MustGet("store").(storage.Store)

	var note *models.Note
	var err error

	req, ok := checkAddUpdNote(c)
	if !ok {
		return
	}

	// Подбираем уникальный id
	i := 0
	id := ""
	for {
		title := req.Title
		id = unidecode.Unidecode(title) + time.Now().Format("_01_02")
		id = utils.Sanitize(id)
		if len(id) < 3 {
			//после очистки длина получилась меньше, пользователь ввел недопустимый title пример "..."
			id = "note"
		}
		if i > 0 {
			id += "_" + strconv.Itoa(i)
		}

		if n, _, _, _ := store.GetNote(id); n == nil {
			break
		}
		i++
	}
	note = &models.Note{
		ID:        id,
		Author:    req.Author,
		Title:     req.Title,
		Password:  req.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Сохраняем заметку
	err = store.UpdateNote(note, req.Content, req.Menu)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save note"})
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

	if config.Cfg.TGBot.MsgOnNewNote {
		link := "https://" + config.Cfg.Site.Host + "/note/" + note.ID
		message := fmt.Sprintf(
			"Создана новая заметка\n\nTitle: %s\n\nLink: %s\n\nID: %s",
			note.Title,
			link,
			note.ID,
		)
		tgbot.SendMessageAll(message)
	}
	c.JSON(http.StatusOK, gin.H{"ok": true, "noteID": note.ID})
}

func EditNote(c *gin.Context) {
	sess := c.MustGet("session").(*models.Session)
	store := c.MustGet("store").(storage.Store)

	var note *models.Note
	var err error

	req, ok := checkAddUpdNote(c)
	if !ok {
		return
	}

	noteID := c.Param("noteID")

	// Смотрим, может ли пользователь редактировать
	hasAccess := false
	for _, id := range sess.Notes {
		if id == noteID {
			hasAccess = true
			break
		}
	}

	// Проверяем пароль
	note, _, _, err = store.GetNote(noteID)
	if err != nil || note == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Note not found"})
		return
	}
	if !hasAccess && note.Password != "" && req.Password != note.Password {
		// Пользователь не может редактировать и в заметке задан пароль и не совпадает
		c.JSON(http.StatusForbidden, gin.H{"error": "wrong password"})
		return
	}

	if !hasAccess && (note.Password == "" || req.Password != note.Password) {
		// Нет в сессии, пароль пустой или не совпадает
		c.JSON(http.StatusForbidden, gin.H{"error": "No access"})
		return
	}

	// Обновляем поля
	note.Title = req.Title
	note.Author = req.Author
	if req.Password != "" {
		note.Password = req.Password
	}
	note.UpdatedAt = time.Now()

	// Сохраняем заметку
	err = store.UpdateNote(note, req.Content, req.Menu)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save note"})
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

	if config.Cfg.TGBot.MsgOnEditNote {
		link := "https://" + config.Cfg.Site.Host + "/note/" + note.ID
		message := fmt.Sprintf(
			"Отредактирована заметка\n\nTitle: %s\n\nLink: %s\n\nID: %s",
			note.Title,
			link,
			note.ID,
		)
		tgbot.SendMessageAll(message)
	}

	c.JSON(http.StatusOK, gin.H{"ok": true, "noteID": note.ID})
}

func isEmptyContent(html string) bool {
	if html == "" {
		return true
	}

	// 1. Если есть хотя бы <img>, <hr>, .ql-formula или <video>, считаем не пустым
	reVisual := regexp.MustCompile(`(?i)<(img|hr|video|a)|class="ql-formula"`)
	if reVisual.MatchString(html) {
		return false
	}

	// 2. Убираем все теги
	reTags := regexp.MustCompile(`<[^>]*>`)
	text := reTags.ReplaceAllString(html, "")

	// 3. Убираем HTML-сущности и пробелы
	text = strings.ReplaceAll(text, "\u00a0", "") // &nbsp;
	text = strings.TrimSpace(text)

	// 4. Если после очистки текста ничего не осталось, контент пустой
	return len(text) == 0
}
