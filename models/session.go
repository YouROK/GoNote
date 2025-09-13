package models

import "time"

type Session struct {
	ID        string    `json:"id"`         // Случайный session_id
	Notes     []string  `json:"notes"`      // Список ID заметок, которые создал пользователь
	ExpiresAt time.Time `json:"expires_at"` // Окончание сессии
}
