package models

import "time"

type UserRole string

const (
	RoleAdmin     UserRole = "admin"
	RoleModerator UserRole = "moderator"
	RoleAuthor    UserRole = "author"
	RoleGuest     UserRole = "guest"
)

type User struct {
	Username  string    `json:"username"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password"` // хеш пароля
	Role      UserRole  `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	Profile   Profile   `json:"profile"`
}

type Profile struct {
	Bio    string `json:"bio,omitempty"`
	Avatar string `json:"avatar,omitempty"`
}
