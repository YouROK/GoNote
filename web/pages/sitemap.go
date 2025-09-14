package pages

import (
	"GoNote/storage/fstorage"
	"encoding/xml"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UrlSet struct {
	XMLName xml.Name `xml:"urlset"`
	Xmlns   string   `xml:"xmlns,attr"`
	Urls    []Url    `xml:"url"`
}

type Url struct {
	Loc        string `xml:"loc"`
	LastMod    string `xml:"lastmod,omitempty"`
	ChangeFreq string `xml:"changefreq,omitempty"`
	Priority   string `xml:"priority,omitempty"`
}

func Sitemap(c *gin.Context) {
	store := c.MustGet("store").(*fstorage.FileStore)

	notes, err := store.ListNotes()
	if err != nil {
		c.String(http.StatusInternalServerError, "Error listing notes: %v", err)
		return
	}

	var urls []Url

	for _, note := range notes {
		urls = append(urls, Url{
			Loc:        "https://gonote.ru/note/" + note.ID,
			LastMod:    note.UpdatedAt.Format("2006-01-02"),
			ChangeFreq: "weekly",
			Priority:   "0.8",
		})
	}

	sitemap := UrlSet{
		Xmlns: "http://www.sitemaps.org/schemas/sitemap/0.9",
		Urls:  urls,
	}

	c.Header("Content-Type", "application/xml")
	c.XML(http.StatusOK, sitemap)
}
