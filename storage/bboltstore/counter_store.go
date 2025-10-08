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
		bucket := tx.Bucket([]byte(CountersBucket))
		if bucket == nil {
			counter = &models.Counter{Count: 0}
			return nil
		}

		counterData := bucket.Get([]byte(noteID))
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
		bucket := tx.Bucket([]byte(CountersBucket))
		if bucket == nil {
			return fmt.Errorf("bucket %s not found", CountersBucket)
		}

		counterData := bucket.Get([]byte(noteID))
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

		return bucket.Put([]byte(noteID), newCounterData)
	})

	return newCounter, err
}
