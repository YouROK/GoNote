package sstorage

import (
	"GoNote/models"
	"database/sql"
	"errors"
)

func (s *SQLiteStore) GetCounterViews(noteID string) (*models.Counter, error) {
	row := s.db.QueryRow(`SELECT count FROM counters WHERE note_id = ?`, noteID)

	var c models.Counter
	if err := row.Scan(&c.Count); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &models.Counter{Count: 0}, nil
		}
		return nil, err
	}

	return &c, nil
}

func (s *SQLiteStore) IncrementCounterViews(noteID string) (*models.Counter, error) {
	// вставляем новую запись или увеличиваем счётчик
	_, err := s.db.Exec(`
		INSERT INTO counters (note_id, count)
		VALUES (?, 1)
		ON CONFLICT(note_id) DO UPDATE SET count = count + 1
	`, noteID)
	if err != nil {
		return nil, err
	}

	return s.GetCounterViews(noteID)
}
