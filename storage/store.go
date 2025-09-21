package storage

import "GoNote/models"

// Интерфейс для работы со страницами
type NoteStore interface {
	CreatePage(uid string, p *models.Note) error
	GetPage(uid, pageID string) (*models.Note, error)
	UpdatePage(uid string, p *models.Note) error
	DeletePage(uid, pageID string) error
	ListPages(uid string) ([]*models.Note, error) // все страницы пользователя
	ListAllPages() ([]*models.Note, error)        // все страницы всех пользователей (для админа)

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

// Интерфейс для работы с счетчиками
type CounterStore interface {
	GetCounterViews(noteID string) (*models.Counter, error)
	IncrementCounterViews(noteID string) (*models.Counter, error)
}
