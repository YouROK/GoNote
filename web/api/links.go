package api

import (
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func GetLinkTitle(c *gin.Context) {
	link := c.Query("url")
	if link == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "url parameter is required"})
		return
	}

	// Валидация: только http/https
	if !strings.HasPrefix(link, "http://") && !strings.HasPrefix(link, "https://") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid URL scheme"})
		return
	}

	// Защита от SSRF: запрещаем локальные адреса
	parsed, err := url.Parse(link)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid URL"})
		return
	}
	if parsed.Hostname() == "localhost" || parsed.Hostname() == "127.0.0.1" || strings.HasPrefix(parsed.Hostname(), "192.168.") || strings.HasPrefix(parsed.Hostname(), "10.") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "local URLs are not allowed"})
		return
	}

	// Выполняем запрос с таймаутом
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	req, err := http.NewRequest("GET", link, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "request failed"})
		return
	}
	req.Header.Set("User-Agent", "GoNote/1.0")

	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": "failed to fetch page"})
		return
	}
	defer resp.Body.Close()

	// Читаем только начало тела (до 10 КБ), чтобы не тратить память
	body, err := io.ReadAll(io.LimitReader(resp.Body, 10240))
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": "failed to read response"})
		return
	}

	// Извлекаем <title> с помощью простого парсинга
	title := extractTitle(string(body))
	if title == "" {
		title = parsed.Hostname() // fallback: имя хоста
	}

	c.JSON(http.StatusOK, gin.H{
		"title": title,
	})
}

func extractTitle(html string) string {
	// 1. Ищем Open Graph title: <meta property="og:title" content="...">
	ogRegex := regexp.MustCompile(`(?i)<meta[^>]+property=["']?og:title["']?[^>]+content=["']([^"']*)["']`)
	if matches := ogRegex.FindStringSubmatch(html); len(matches) > 1 {
		title := strings.TrimSpace(matches[1])
		if title != "" {
			return cleanTitle(title)
		}
	}

	// 2. Ищем Twitter title: <meta name="twitter:title" content="...">
	twitterRegex := regexp.MustCompile(`(?i)<meta[^>]+name=["']?twitter:title["']?[^>]+content=["']([^"']*)["']`)
	if matches := twitterRegex.FindStringSubmatch(html); len(matches) > 1 {
		title := strings.TrimSpace(matches[1])
		if title != "" {
			return cleanTitle(title)
		}
	}

	// 3. Падаем обратно на <title>...</title>
	titleRegex := regexp.MustCompile(`(?i)<title[^>]*>([^<]*)</title>`)
	if matches := titleRegex.FindStringSubmatch(html); len(matches) > 1 {
		title := strings.TrimSpace(matches[1])
		if title != "" {
			return cleanTitle(title)
		}
	}

	return ""
}

// Вспомогательная функция для очистки (можно расширить)
func cleanTitle(title string) string {
	// Удаляем HTML-сущности (минимум)
	title = strings.ReplaceAll(title, "&nbsp;", " ")
	title = strings.ReplaceAll(title, "&#39;", "'")
	title = strings.ReplaceAll(title, "&quot;", `"`)
	title = strings.ReplaceAll(title, "&amp;", "&")
	return title
}
