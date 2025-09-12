package storage

import "GoNote/models"

// Интерфейс для работы с пользователями
//type UserStore interface {
//	CreateUser(u *models.User) error
//	GetUser(uid string) (*models.User, error)
//	UpdateUser(u *models.User) error
//	DeleteUser(uid string) error
//	UserExists(uid string) bool
//	ListUsers() ([]*models.User, error) // для админки
//}

// Интерфейс для работы со страницами
type PageStore interface {
	CreatePage(uid string, p *models.Page) error
	GetPage(uid, pageID string) (*models.Page, error)
	UpdatePage(uid string, p *models.Page) error
	DeletePage(uid, pageID string) error
	ListPages(uid string) ([]*models.Page, error) // все страницы пользователя
	ListAllPages() ([]*models.Page, error)        // все страницы всех пользователей (для админа)

	SavePageContent(uid, pageID, content string) error
	LoadPageContent(uid, pageID string) (string, error)

	SaveAsset(uid, pageID, filename string, data []byte) error
	LoadAsset(uid, pageID, filename string) ([]byte, error)
	ListAssets(uid, pageID string) ([]string, error)
	DeleteAsset(uid, pageID, filename string) error
}

// Интерфейс для работы с сессиями
type SessionStore interface {
	CreateSession(uid string) (string, error)
	DeleteSession(uid string) error
	GetSession(uid string) (string, error)

	LoadSession(uid string) (string, error)
	SaveSession(uid, token string) error
}
