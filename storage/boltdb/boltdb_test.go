package boltdb

import (
	"os"
	"testing"

	"github.com/DavidHuie/beef/Godeps/_workspace/src/github.com/boltdb/bolt"
)

func TestBits(t *testing.T) {
	// So that testing traverses several pages
	SetPageSize(1)

	db, err := bolt.Open("db.test", 0600, nil)
	if err != nil {
		panic(err)
	}

	n := uint64(1234567801234567890)

	// Insert many bits into the bucket
	err = db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucket([]byte("test"))
		if err != nil {
			t.Error(err)
		}

		storage := New(b)

		for i := uint64(0); i < 64; i++ {
			mask := n & (uint64(1) << i)
			val := (n & mask) >> i
			if val == 1 {
				storage.SetBit(i)
			}
		}

		return nil
	})
	if err != nil {
		t.Error(err)
	}

	// Check if all bits were set
	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("test"))
		storage := New(b)

		for i := uint64(0); i < 64; i++ {
			mask := n & (uint64(1) << i)
			val := (n & mask) >> i

			bit, err := storage.GetBit(i)
			if err != nil {
				t.Error(err)
			}

			if val == 1 && !bit {
				t.Errorf("bit not set correctly")
			}
			if val == 0 && bit {
				t.Errorf("bit not set correctly")
			}
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
