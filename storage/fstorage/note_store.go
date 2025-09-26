package fstorage

import (
	"GoNote/models"
	"encoding/json"
	"os"
	"path/filepath"
)

func (fs *FileStore) CreateNote(n *models.Note, content, menu string) error {
	notePath := filepath.Join(fs.notes, n.ID)
	if err := os.MkdirAll(notePath, 0755); err != nil {
		return err
	}

	// Сохраняем JSON с метаданными
	metaPath := filepath.Join(notePath, "note.json")
	data, _ := json.MarshalIndent(n, "", "  ")
	if err := os.WriteFile(metaPath, data, 0644); err != nil {
		os.RemoveAll(notePath)
		return err
	}

	// Сохраняем контент
	err := os.WriteFile(filepath.Join(notePath, "content.md"), []byte(content), 0644)
	if err != nil {
		os.RemoveAll(notePath)
		return err
	}

	// Сохраняем меню
	err = os.WriteFile(filepath.Join(notePath, "menu.md"), []byte(menu), 0644)
	if err != nil {
		os.RemoveAll(notePath)
		return err
	}
	return nil
}

func (fs *FileStore) GetNote(noteID string) (*models.Note, string, string, error) {
	notePath := filepath.Join(fs.notes, noteID)

	// Читаем JSON
	metaPath := filepath.Join(notePath, "note.json")
	data, err := os.ReadFile(metaPath)
	if err != nil {
		return nil, "", "", err
	}
	var n models.Note
	if err := json.Unmarshal(data, &n); err != nil {
		return nil, "", "", err
	}

	// Читаем content.md
	contentPath := filepath.Join(notePath, "content.md")
	content, err := os.ReadFile(contentPath)
	if err != nil {
		return nil, "", "", err
	}

	// Читаем menu.md
	menuPath := filepath.Join(notePath, "menu.md")
	menu, err := os.ReadFile(menuPath)
	if err != nil && !os.IsNotExist(err) {
		return nil, "", "", err
	}

	return &n, string(content), string(menu), nil
}

func (fs *FileStore) UpdateNote(n *models.Note, content, menu string) error {
	return fs.CreateNote(n, content, menu) // перезаписываем
}

func (fs *FileStore) DeleteNote(noteID string) error {
	return os.RemoveAll(filepath.Join(fs.notes, noteID))
}

func (fs *FileStore) ListNotes() ([]*models.Note, error) {
	entries, err := os.ReadDir(fs.notes)
	if err != nil {
		return nil, err
	}

	var notes []*models.Note
	for _, e := range entries {
		if !e.IsDir() {
			continue
		}
		n, _, _, err := fs.GetNote(e.Name())
		if err == nil {
			notes = append(notes, n)
		}
	}
	return notes, nil
}
