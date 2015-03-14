package bloomfilter

import (
	"hash/fnv"
	"os"
	"testing"

	"github.com/DavidHuie/beef/Godeps/_workspace/src/github.com/boltdb/bolt"
	"github.com/DavidHuie/beef/storage/boltdb"
)

func TestInsertAndCheck(t *testing.T) {
	db, err := bolt.Open("db.test", 0600, nil)
	if err != nil {
		panic(err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucket([]byte("test"))
		if err != nil {
			t.Error(err)
		}

		storage := boltdb.New(b)

		bf := New(storage, fnv.New64(), 3, 1000000)

		val, err := bf.Check([]byte("hi"))
		if err != nil {
			return err
		}
		if val {
			t.Errorf("entry should not exist")
		}

		if err := bf.Insert([]byte("hi")); err != nil {
			t.Error(err)
		}

		val, err = bf.Check([]byte("hi"))
		if err != nil {
			return err
		}
		if !val {
			t.Errorf("entry should exist")
		}

		return nil
	})
	if err != nil {
		t.Error(err)
	}

	if err := os.Remove("db.test"); err != nil {
		t.Error(err)
	}
}
