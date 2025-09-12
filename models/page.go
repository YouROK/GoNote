package models

import "time"

type NoteStatus string

const (
	StatusDraft     NoteStatus = "draft"
	StatusPublished NoteStatus = "published"
	StatusHidden    NoteStatus = "hidden"
)

type Note struct {
	ID        string     `json:"id"`     // хэш страницы, уникальный
	Author    string     `json:"author"` // uid или имя пользователя
	Title     string     `json:"title"`
	Password  string     `json:"password"` // bcrypt-хеш, если задан
	Status    NoteStatus `json:"status"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}
