package sstorage

import (
	"database/sql"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

type SQLiteStore struct {
	db *sql.DB
}

func NewSQLiteStore(dbDir string) (*SQLiteStore, error) {
	fname := filepath.Join(dbDir, "gonote.db")
	db, err := sql.Open("sqlite3", fname)
	if err != nil {
		return nil, err
	}

	s := &SQLiteStore{db: db}
	if err := s.initSchema(); err != nil {
		return nil, err
	}

	return s, nil
}

func (s *SQLiteStore) initSchema() error {
	_, err := s.db.Exec(`
		CREATE TABLE IF NOT EXISTS notes (
			id TEXT PRIMARY KEY,
			author TEXT,
			title TEXT,
			password TEXT,
			content TEXT,
			created_at TEXT,
			updated_at TEXT
		);

		CREATE TABLE IF NOT EXISTS sessions (
			id TEXT PRIMARY KEY,
			notes TEXT,
			expires_at TEXT
		);

		CREATE TABLE IF NOT EXISTS counters (
			note_id TEXT PRIMARY KEY,
			count INTEGER
		);
	`)
	return err
}
