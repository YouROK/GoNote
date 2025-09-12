package fstorage

import (
	"GoNote/models"
	"encoding/json"
	"os"
	"path/filepath"
)

func (fs *FileStore) CreateNote(n *models.Note, content string) error {
	notePath := filepath.Join(fs.notes, n.ID)
	if err := os.MkdirAll(notePath, 0755); err != nil {
		return err
	}

	// Сохраняем JSON с метаданными
	metaPath := filepath.Join(notePath, "note.json")
	data, _ := json.MarshalIndent(n, "", "  ")
	if err := os.WriteFile(metaPath, data, 0644); err != nil {
		return err
	}

	// Сохраняем контент
	return os.WriteFile(filepath.Join(notePath, "content.md"), []byte(content), 0644)
}

func (fs *FileStore) GetNote(noteID string) (*models.Note, string, error) {
	notePath := filepath.Join(fs.notes, noteID)

	// Читаем JSON
	metaPath := filepath.Join(notePath, "note.json")
	data, err := os.ReadFile(metaPath)
	if err != nil {
		return nil, "", err
	}
	var n models.Note
	if err := json.Unmarshal(data, &n); err != nil {
		return nil, "", err
	}

	// Читаем content.md
	contentPath := filepath.Join(notePath, "content.md")
	content, err := os.ReadFile(contentPath)
	if err != nil {
		return nil, "", err
	}

	return &n, string(content), nil
}

func (fs *FileStore) UpdateNote(n *models.Note, content string) error {
	return fs.CreateNote(n, content) // перезаписываем
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
		n, _, err := fs.GetNote(e.Name())
		if err == nil {
			notes = append(notes, n)
		}
	}
	return notes, nil
}

// ===== Работа с assets =====

func (fs *FileStore) SaveAsset(noteID, filename string, data []byte) error {
	assetsPath := filepath.Join(fs.notes, noteID, "assets")
	if err := os.MkdirAll(assetsPath, 0755); err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(assetsPath, filename), data, 0644)
}

func (fs *FileStore) LoadAsset(noteID, filename string) ([]byte, error) {
	return os.ReadFile(filepath.Join(fs.notes, noteID, "assets", filename))
}

func (fs *FileStore) ListAssets(noteID string) ([]string, error) {
	assetsPath := filepath.Join(fs.notes, noteID, "assets")
	files, err := os.ReadDir(assetsPath)
	if err != nil {
		return nil, err
	}
	var list []string
	for _, f := range files {
		if !f.IsDir() {
			list = append(list, f.Name())
		}
	}
	return list, nil
}

func (fs *FileStore) DeleteAsset(noteID, filename string) error {
	return os.Remove(filepath.Join(fs.notes, noteID, "assets", filename))
}
