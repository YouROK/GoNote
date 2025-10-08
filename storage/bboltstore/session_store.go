package bboltstore

import (
	"GoNote/models"
	"GoNote/utils"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"go.etcd.io/bbolt"
)

func (bs *BboltStore) CreateSession() (*models.Session, error) {
	sessionID := utils.RandStr(16) // Импортируйте utils
	sess := &models.Session{
		ID:        sessionID,
		Notes:     []string{},
		ExpiresAt: time.Now().Add(7 * 24 * time.Hour),
	}

	err := bs.db.Update(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(SessionsBucket))
		if bucket == nil {
			return fmt.Errorf("bucket %s not found", SessionsBucket)
		}

		sessionData, err := json.Marshal(sess)
		if err != nil {
			return err
		}

		return bucket.Put([]byte(sessionID), sessionData)
	})

	if err != nil {
		return nil, err
	}

	return sess, nil
}

func (bs *BboltStore) LoadSession(sessionID string) (*models.Session, error) {
	var sess *models.Session
	var err error

	err = bs.db.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(SessionsBucket))
		if bucket == nil {
			return fmt.Errorf("bucket %s not found", SessionsBucket)
		}

		sessionData := bucket.Get([]byte(sessionID))
		if sessionData == nil {
			return fmt.Errorf("session with id %s not found", sessionID)
		}

		sess = &models.Session{}
		if err := json.Unmarshal(sessionData, sess); err != nil {
			return err
		}

		if time.Now().After(sess.ExpiresAt) {
			bs.DeleteSession(sessionID)
			return fmt.Errorf("session expired")
		}

		return nil
	})

	return sess, err
}

func (bs *BboltStore) SaveSession(sess *models.Session) error {
	return bs.db.Update(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(SessionsBucket))
		if bucket == nil {
			return fmt.Errorf("bucket %s not found", SessionsBucket)
		}

		sessionData, err := json.Marshal(sess)
		if err != nil {
			return err
		}

		return bucket.Put([]byte(sess.ID), sessionData)
	})
}

func (bs *BboltStore) DeleteSession(sessionID string) error {
	return bs.db.Update(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(SessionsBucket))
		if bucket == nil {
			return fmt.Errorf("bucket %s not found", SessionsBucket)
		}

		return bucket.Delete([]byte(sessionID))
	})
}

func (bs *BboltStore) SessionExists(sessionID string) bool {
	var exists bool
	bs.db.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(SessionsBucket))
		if bucket == nil {
			exists = false
			return nil
		}
		sessionData := bucket.Get([]byte(sessionID))
		exists = sessionData != nil
		return nil
	})
	return exists
}

func (bs *BboltStore) RemoveExpiredSessions() {
	err := bs.db.Update(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(SessionsBucket))
		if bucket == nil {
			return fmt.Errorf("bucket %s not found", SessionsBucket)
		}

		return bucket.ForEach(func(k, v []byte) error {
			var ses *models.Session
			err := json.Unmarshal(v, &ses)
			if err != nil {
				return err
			}
			if time.Now().After(ses.ExpiresAt) {
				return bucket.Delete(k)
			}
			return nil
		})
	})

	if err != nil {
		log.Println("Error removing expired sessions: ", err)
	}
}
