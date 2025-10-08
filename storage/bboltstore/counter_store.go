package bboltstore

import (
	"GoNote/models"
	"encoding/json"
	"fmt"

	"go.etcd.io/bbolt"
)

func (bs *BboltStore) GetCounterViews(noteID string) (*models.Counter, error) {
	var counter *models.Counter
	var err error

	err = bs.db.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(NotesBucket))
		if bucket == nil {
			counter = &models.Counter{Count: 0}
			return nil
		}

		noteBucket := bucket.Bucket([]byte(noteID))
		if bucket == nil {
			counter = &models.Counter{Count: 0}
			return nil
		}

		counterData := noteBucket.Get([]byte("counter"))
		if counterData == nil {
			counter = &models.Counter{Count: 0}
			return nil
		}

		counter = &models.Counter{}
		if err := json.Unmarshal(counterData, counter); err != nil {
			counter = &models.Counter{Count: 0}
			return nil
		}

		return nil
	})

	return counter, err
}

func (bs *BboltStore) IncrementCounterViews(noteID string) (*models.Counter, error) {
	var newCounter *models.Counter
	var err error

	err = bs.db.Update(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(NotesBucket))
		if bucket == nil {
			return fmt.Errorf("bucket %s not found", NotesBucket)
		}

		noteBucket := bucket.Bucket([]byte(noteID))
		if bucket == nil {
			return fmt.Errorf("%s not found", noteID)
		}

		counterData := noteBucket.Get([]byte("counter"))
		var currentCounter models.Counter
		if counterData != nil {
			if err := json.Unmarshal(counterData, &currentCounter); err != nil {
				currentCounter.Count = 0
			}
		}

		currentCounter.Count++

		newCounter = &currentCounter
		newCounterData, err := json.Marshal(newCounter)
		if err != nil {
			return err
		}

		return noteBucket.Put([]byte("counter"), newCounterData)
	})

	return newCounter, err
}
