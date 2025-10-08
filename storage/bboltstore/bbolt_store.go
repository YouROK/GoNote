package bboltstore

import (
	"os"
	"path/filepath"

	"go.etcd.io/bbolt"
)

const (
	NotesBucket    = "notes"
	SessionsBucket = "sessions"
	CountersBucket = "counters"
)

type BboltStore struct {
	dbFile string
	db     *bbolt.DB
}

func NewBboltStore(dbDir string) (*BboltStore, error) {
	if err := os.MkdirAll(dbDir, 0755); err != nil {
		return nil, err
	}

	db, err := bbolt.Open(filepath.Join(dbDir, "gonote.bdb"), 0600, nil)
	if err != nil {
		return nil, err
	}

	err = db.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(NotesBucket))
		if err != nil {
			return err
		}
		_, err = tx.CreateBucketIfNotExists([]byte(SessionsBucket))
		if err != nil {
			return err
		}
		_, err = tx.CreateBucketIfNotExists([]byte(CountersBucket))
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		db.Close()
		return nil, err
	}

	return &BboltStore{dbFile: dbDir, db: db}, nil
}

func (bs *BboltStore) Close() error {
	if bs.db != nil {
		err := bs.db.Close()
		bs.db = nil
		return err
	}
	return nil
}
