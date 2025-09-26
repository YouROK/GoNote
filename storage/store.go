package storage

import (
	"GoNote/models"
	"GoNote/storage/fstorage"
	"errors"
)

const FS_STORE = 0
const SQLITE_STORE = 1

func NewStore(typeStor int, dir string) (Store, error) {
	if typeStor == FS_STORE {
		return fstorage.NewFileStore(dir), nil
		//} else {
		//	return sstorage.NewSQLiteStore(dir)
	}
	return nil, errors.New("type store not support")
}

type Store interface {
	// ---- Работа с заметками ----
	CreateNote(n *models.Note, content, menu string) error
	GetNote(noteID string) (*models.Note, string, string, error)
	UpdateNote(n *models.Note, content, menu string) error
	DeleteNote(noteID string) error
	ListNotes() ([]*models.Note, error)

	// ---- Работа с сессиями ----
	CreateSession() (*models.Session, error)
	LoadSession(sessionID string) (*models.Session, error)
	SaveSession(sess *models.Session) error
	DeleteSession(sessionID string) error
	SessionExists(sessionID string) bool

	// ---- Работа со счетчиком просмотров ----
	GetCounterViews(noteID string) (*models.Counter, error)
	IncrementCounterViews(noteID string) (*models.Counter, error)
}
