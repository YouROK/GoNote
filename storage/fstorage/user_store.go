package fstorage

import (
	"GoNote/models"
	"encoding/json"
	"os"
	"path/filepath"
)

// ===== UserStore методы =====

func (fs *FileStore) CreateUser(u *models.User) error {
	userPath := filepath.Join(fs.users, u.Username)
	if _, err := os.Stat(userPath); !os.IsNotExist(err) {
		return os.ErrExist
	}
	if err := os.MkdirAll(userPath, 0755); err != nil {
		return err
	}
	data, _ := json.MarshalIndent(u, "", "  ")
	return os.WriteFile(filepath.Join(userPath, "user.json"), data, 0644)
}

func (fs *FileStore) GetUser(username string) (*models.User, error) {
	userPath := filepath.Join(fs.users, username, "user.json")
	data, err := os.ReadFile(userPath)
	if err != nil {
		return nil, err
	}
	var u models.User
	if err := json.Unmarshal(data, &u); err != nil {
		return nil, err
	}
	return &u, nil
}

func (fs *FileStore) UpdateUser(u *models.User) error {
	return fs.CreateUser(u) // перезаписываем
}

func (fs *FileStore) DeleteUser(username string) error {
	return os.RemoveAll(filepath.Join(fs.users, username))
}

func (fs *FileStore) UserExists(username string) bool {
	info, err := os.Stat(filepath.Join(fs.users, username))
	return err == nil && info.IsDir()
}

func (fs *FileStore) ListUsers() ([]*models.User, error) {
	entries, err := os.ReadDir(fs.users)
	if err != nil {
		return nil, err
	}
	var users []*models.User
	for _, e := range entries {
		if !e.IsDir() {
			continue
		}
		u, err := fs.GetUser(e.Name())
		if err == nil {
			users = append(users, u)
		}
	}
	return users, nil
}
