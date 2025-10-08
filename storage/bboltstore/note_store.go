package bboltstore

import (
	"GoNote/models"
	"encoding/json"
	"fmt"
	"log"

	"go.etcd.io/bbolt"
)

func (bs *BboltStore) CreateNote(n *models.Note, content, menu string) error {
	return bs.db.Update(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(NotesBucket))
		if bucket == nil {
			return fmt.Errorf("%s not found", NotesBucket)
		}

		noteBucket, err := bucket.CreateBucketIfNotExists([]byte(n.ID))
		if err != nil {
			return err
		}

		noteData, err := json.Marshal(n)
		if err != nil {
			return err
		}

		if err := noteBucket.Put([]byte("note"), noteData); err != nil {
			return err
		}
		if err := noteBucket.Put([]byte("content"), []byte(content)); err != nil {
			return err
		}
		if err := noteBucket.Put([]byte("menu"), []byte(menu)); err != nil {
			return err
		}

		return nil
	})
}

func (bs *BboltStore) GetNote(noteID string) (*models.Note, string, string, error) {
	var note *models.Note
	var content, menu string

	err := bs.db.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(NotesBucket))
		if bucket == nil {
			return fmt.Errorf("bucket %s not found", NotesBucket)
		}

		noteBucket := bucket.Bucket([]byte(noteID))
		if noteBucket == nil {
			return fmt.Errorf("%s not found", noteID)
		}

		noteData := noteBucket.Get([]byte("note"))
		if noteData == nil {
			return fmt.Errorf("note with id %s not found", noteID)
		}

		note = &models.Note{}
		if err := json.Unmarshal(noteData, note); err != nil {
			return err
		}

		contentBytes := noteBucket.Get([]byte("content"))
		menuBytes := noteBucket.Get([]byte("menu"))

		if contentBytes != nil {
			content = string(contentBytes)
		} else {
			content = ""
		}
		if menuBytes != nil {
			menu = string(menuBytes)
		} else {
			menu = ""
		}

		return nil
	})

	return note, content, menu, err
}

func (bs *BboltStore) UpdateNote(n *models.Note, content, menu string) error {
	return bs.CreateNote(n, content, menu)
}

func (bs *BboltStore) DeleteNote(noteID string) error {
	return bs.db.Update(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(NotesBucket))
		if bucket == nil {
			return fmt.Errorf("%s not found", NotesBucket)
		}

		if err := bucket.DeleteBucket([]byte(noteID)); err != nil {
			return err
		}

		return nil
	})
}

func (bs *BboltStore) ListNotes() ([]*models.Note, error) {
	var notes []*models.Note

	err := bs.db.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(NotesBucket))
		if bucket == nil {
			return fmt.Errorf("%s not found", NotesBucket)
		}

		return bucket.ForEachBucket(func(k []byte) error {
			if k != nil {
				noteBucket := bucket.Bucket(k)
				if noteBucket != nil {
					noteData := noteBucket.Get([]byte("note"))
					var note models.Note
					if err := json.Unmarshal(noteData, &note); err != nil {
						log.Printf("Error unmarshaling note %s: %v", string(k), err)
						return nil
					}
					notes = append(notes, &note)
				}
			}
			return nil
		})
	})

	return notes, err
}
