package sstorage

import (
	"database/sql"
	"errors"
	"time"

	"GoNote/models"
)

// CreateNote сохраняет новую заметку и её контент
func (s *SQLiteStore) CreateNote(n *models.Note, content string) error {
	_, err := s.db.Exec(`
		INSERT INTO notes (id, author, title, password, content, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?)
	`, n.ID, n.Author, n.Title, n.Password, content, n.CreatedAt.Format(time.RFC3339Nano), n.UpdatedAt.Format(time.RFC3339Nano))
	return err
}

// GetNote загружает заметку и её контент по ID
func (s *SQLiteStore) GetNote(noteID string) (*models.Note, string, error) {
	row := s.db.QueryRow(`
		SELECT id, author, title, password, content, created_at, updated_at
		FROM notes
		WHERE id = ?
	`, noteID)

	var n models.Note
	var content, createdAt, updatedAt string

	if err := row.Scan(&n.ID, &n.Author, &n.Title, &n.Password, &content, &createdAt, &updatedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, "", nil
		}
		return nil, "", err
	}

	n.CreatedAt, _ = time.Parse(time.RFC3339Nano, createdAt)
	n.UpdatedAt, _ = time.Parse(time.RFC3339Nano, updatedAt)

	return &n, content, nil
}

// UpdateNote обновляет заметку и её контент
func (s *SQLiteStore) UpdateNote(n *models.Note, content string) error {
	_, err := s.db.Exec(`
		INSERT INTO notes (id, author, title, password, content, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?)
		ON CONFLICT(id) DO UPDATE SET
			author = excluded.author,
			title = excluded.title,
			password = excluded.password,
			content = excluded.content,
			updated_at = excluded.updated_at
	`, n.ID, n.Author, n.Title, n.Password, content, n.CreatedAt.Format(time.RFC3339Nano), n.UpdatedAt.Format(time.RFC3339Nano))
	return err
}

// DeleteNote удаляет заметку по ID
func (s *SQLiteStore) DeleteNote(noteID string) error {
	_, err := s.db.Exec(`DELETE FROM notes WHERE id = ?`, noteID)
	return err
}

// ListNotes возвращает список заметок без контента
func (s *SQLiteStore) ListNotes() ([]*models.Note, error) {
	rows, err := s.db.Query(`
		SELECT id, author, title, password, created_at, updated_at
		FROM notes
		ORDER BY created_at DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notes []*models.Note
	for rows.Next() {
		var n models.Note
		var createdAt, updatedAt string

		if err := rows.Scan(&n.ID, &n.Author, &n.Title, &n.Password, &createdAt, &updatedAt); err != nil {
			return nil, err
		}

		n.CreatedAt, _ = time.Parse(time.RFC3339Nano, createdAt)
		n.UpdatedAt, _ = time.Parse(time.RFC3339Nano, updatedAt)

		notes = append(notes, &n)
	}
	return notes, nil
}
