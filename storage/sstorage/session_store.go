package sstorage

import (
	"GoNote/models"
	"GoNote/utils"
	"database/sql"
	"errors"
	"time"
)

// Создание новой сессии
func (s *SQLiteStore) CreateSession() (*models.Session, error) {
	sessionID := utils.RandStr(16)
	expires := time.Now().Add(7 * 24 * time.Hour)

	_, err := s.db.Exec(`INSERT INTO sessions (id, expires_at) VALUES (?, ?)`, sessionID, expires)
	if err != nil {
		return nil, err
	}

	return &models.Session{
		ID:        sessionID,
		Notes:     []string{},
		ExpiresAt: expires,
	}, nil
}

// Загрузка сессии по ID
func (s *SQLiteStore) LoadSession(sessionID string) (*models.Session, error) {
	row := s.db.QueryRow(`SELECT expires_at FROM sessions WHERE id = ?`, sessionID)
	var expires time.Time
	if err := row.Scan(&expires); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, err
	}

	if time.Now().After(expires) {
		_ = s.DeleteSession(sessionID)
		return nil, sql.ErrNoRows
	}

	// грузим привязанные заметки
	rows, err := s.db.Query(`SELECT note_id FROM session_notes WHERE session_id = ?`, sessionID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notes []string
	for rows.Next() {
		var nid string
		if err := rows.Scan(&nid); err == nil {
			notes = append(notes, nid)
		}
	}

	return &models.Session{
		ID:        sessionID,
		Notes:     notes,
		ExpiresAt: expires,
	}, nil
}

// Сохраняем (обновляем) сессию
func (s *SQLiteStore) SaveSession(sess *models.Session) error {
	_, err := s.db.Exec(
		`INSERT INTO sessions (id, expires_at) VALUES (?, ?)
		 ON CONFLICT(id) DO UPDATE SET expires_at = excluded.expires_at`,
		sess.ID, sess.ExpiresAt,
	)
	if err != nil {
		return err
	}

	// сначала чистим записи
	_, err = s.db.Exec(`DELETE FROM session_notes WHERE session_id = ?`, sess.ID)
	if err != nil {
		return err
	}

	// потом вставляем новые
	for _, noteID := range sess.Notes {
		_, err = s.db.Exec(`INSERT INTO session_notes (session_id, note_id) VALUES (?, ?)`, sess.ID, noteID)
		if err != nil {
			return err
		}
	}

	return nil
}

// Удаляем сессию
func (s *SQLiteStore) DeleteSession(sessionID string) error {
	_, err := s.db.Exec(`DELETE FROM session_notes WHERE session_id = ?`, sessionID)
	if err != nil {
		return err
	}
	_, err = s.db.Exec(`DELETE FROM sessions WHERE id = ?`, sessionID)
	return err
}

// Проверяем, существует ли сессия
func (s *SQLiteStore) SessionExists(sessionID string) bool {
	row := s.db.QueryRow(`SELECT 1 FROM sessions WHERE id = ? LIMIT 1`, sessionID)
	var tmp int
	if err := row.Scan(&tmp); err != nil {
		return false
	}
	return true
}
