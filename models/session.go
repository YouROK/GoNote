package models

type Session struct {
	ID    string   `json:"id"`    // случайный session_id
	Notes []string `json:"notes"` // список ID заметок, которые создал пользователь
}
