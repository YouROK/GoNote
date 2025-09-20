package fstorage

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Counter struct {
	Count int `json:"count"`
}

// GetCounter возвращает текущее значение счетчика заметки
func (fs *FileStore) GetCounter(noteID string) (*Counter, error) {
	counterPath := filepath.Join(fs.notes, noteID, "counter.json")

	data, err := os.ReadFile(counterPath)
	if err != nil {
		if os.IsNotExist(err) {
			return &Counter{}, nil // если файла нет, считаем 0
		}
		return nil, err
	}

	var c *Counter
	if err := json.Unmarshal(data, &c); err != nil {
		return &Counter{}, nil // если файл битый, начинаем с 0
	}

	return c, nil
}

// IncrementCounter увеличивает счетчик на 1 и возвращает новое значение
func (fs *FileStore) IncrementCounter(noteID string) (*Counter, error) {
	c, _ := fs.GetCounter(noteID)
	if c == nil {
		c = &Counter{}
	}
	c.Count++

	counterPath := filepath.Join(fs.notes, noteID, "counter.json")
	data, _ := json.MarshalIndent(c, "", "  ")

	if err := os.WriteFile(counterPath, data, 0644); err != nil {
		return &Counter{}, err
	}

	return c, nil
}
