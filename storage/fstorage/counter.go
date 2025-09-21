package fstorage

import (
	"GoNote/models"
	"encoding/json"
	"os"
	"path/filepath"
)

// GetCounterViews возвращает текущее значение счетчика заметки
func (fs *FileStore) GetCounterViews(noteID string) (*models.Counter, error) {
	counterPath := filepath.Join(fs.notes, noteID, "counter.json")

	data, err := os.ReadFile(counterPath)
	if err != nil {
		if os.IsNotExist(err) {
			return &models.Counter{}, nil
		}
		return nil, err
	}

	var c *models.Counter
	if err := json.Unmarshal(data, &c); err != nil {
		return &models.Counter{}, nil
	}

	return c, nil
}

// IncrementCounterViews увеличивает счетчик на 1 и возвращает новое значение
func (fs *FileStore) IncrementCounterViews(noteID string) (*models.Counter, error) {
	c, _ := fs.GetCounterViews(noteID)
	if c == nil {
		c = &models.Counter{}
	}
	c.Count++

	counterPath := filepath.Join(fs.notes, noteID, "counter.json")
	data, _ := json.MarshalIndent(c, "", "  ")

	if err := os.WriteFile(counterPath, data, 0644); err != nil {
		return &models.Counter{}, err
	}

	return c, nil
}
