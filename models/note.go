package models

import "time"

type Note struct {
	ID        string    `json:"id"`
	Author    string    `json:"author"`
	Title     string    `json:"title"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
