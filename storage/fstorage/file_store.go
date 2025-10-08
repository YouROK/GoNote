package fstorage

import "path/filepath"

type FileStore struct {
	DBDir    string
	notes    string
	sessions string
	users    string
}

func NewFileStore(dbDir string) *FileStore {
	return &FileStore{
		DBDir:    dbDir,
		notes:    filepath.Join(dbDir, "notes"),
		sessions: filepath.Join(dbDir, "sessions"),
		users:    filepath.Join(dbDir, "users"),
	}
}

func (fs *FileStore) Close() error {
	return nil
}
