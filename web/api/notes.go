package api

import (
	"GoNote/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetNoteApi(c *gin.Context) {
	store := c.MustGet("store").(storage.Store)

	noteID := c.Param("noteID")

	note, _, _, err := store.GetNote(noteID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "note not found",
		})
		return
	}

	hasPass := note.Password != ""

	counter, _ := store.GetCounterViews(noteID)

	c.JSON(http.StatusOK, gin.H{
		"author":     note.Author,
		"title":      note.Title,
		"created_at": note.CreatedAt,
		"updated_at": note.UpdatedAt,
		"has_pass":   hasPass,
		"view_count": counter.Count,
	})
}

func GetNoteContentApi(c *gin.Context) {
	store := c.MustGet("store").(storage.Store)

	noteID := c.Param("noteID")

	_, content, _, err := store.GetNote(noteID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "note not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"content": content,
	})
}

func GetNoteMenuApi(c *gin.Context) {
	store := c.MustGet("store").(storage.Store)

	noteID := c.Param("noteID")

	_, _, menu, err := store.GetNote(noteID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "note not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"menu": menu,
	})
}
