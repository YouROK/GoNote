package fstorage

import (
	"GoNote/models"
	"GoNote/utils"
	"encoding/json"
	"os"
	"path/filepath"
	"time"
)

func (fs *FileStore) sessionPath(sessionID string) string {
	return filepath.Join(fs.sessions, sessionID+".json")
}

// Создание новой сессии
func (fs *FileStore) CreateSession() (*models.Session, error) {
	sessionID := utils.RandStr(16)

	sess := &models.Session{
		ID:        sessionID,
		Notes:     []string{},
		ExpiresAt: time.Now().Add(7 * 24 * time.Hour), // сессия живёт 7 дней
	}

	if err := os.MkdirAll(fs.sessions, 0755); err != nil {
		return nil, err
	}

	data, _ := json.MarshalIndent(sess, "", "  ")
	if err := os.WriteFile(fs.sessionPath(sessionID), data, 0644); err != nil {
		return nil, err
	}

	return sess, nil
}

// Загружаем сессию по ID
func (fs *FileStore) LoadSession(sessionID string) (*models.Session, error) {
	data, err := os.ReadFile(fs.sessionPath(sessionID))
	if err != nil {
		return nil, err
	}

	var sess models.Session
	if err := json.Unmarshal(data, &sess); err != nil {
		return nil, err
	}

	if time.Now().After(sess.ExpiresAt) {
		fs.DeleteSession(sessionID)
		return nil, os.ErrNotExist
	}

	return &sess, nil
}

// Сохраняем сессию
func (fs *FileStore) SaveSession(sess *models.Session) error {
	data, _ := json.Marshal(sess)
	return os.WriteFile(fs.sessionPath(sess.ID), data, 0644)
}

// Удаляем сессию
func (fs *FileStore) DeleteSession(sessionID string) error {
	return os.Remove(fs.sessionPath(sessionID))
}

// Добавляем заметку в сессию
func (fs *FileStore) AddNoteToSession(sessionID, noteID string) error {
	sess, err := fs.LoadSession(sessionID)
	if err != nil {
		return err
	}

	sess.Notes = append(sess.Notes, noteID)
	return fs.SaveSession(sess)
}

// Проверяем, существует ли сессия
func (fs *FileStore) SessionExists(sessionID string) bool {
	_, err := os.Stat(fs.sessionPath(sessionID))
	return err == nil
}
