package beef

import (
	"errors"

	"github.com/DavidHuie/beef/Godeps/_workspace/src/github.com/boltdb/bolt"
)

type DB struct {
	db *bolt.DB
}

const (
	// The number of bytes to store in each key
	pageSize = 8
)

var (
	ErrBFMissing       = errors.New("bloom filter does not exist")
	ErrMetadataMissing = errors.New("metadata is missing")
	ErrBFAlreadyExists = errors.New("bloom filter already exists")

	metadataDbKey = []byte("m")
)

func (db *DB) GetMetadata(tx *bolt.Tx, name string) (*Metadata, error) {
	// err := db.db.View(func(tx *bolt.Tx) error {
	// })
	return nil, nil
}

func (db *DB) CreateBloomFilter(name string, hashFunction, numHashFunctions, size uint64) error {
	// return db.db.Update(func(tx *bolt.Tx) error {
	// })
	return nil
}

func (db *DB) BFInsert(bfName string, value []byte) error {
	// return db.db.Update(func(tx *bolt.Tx) error {
	// 	b := tx.Bucket(bucketNameForBF(name))
	// 	if b == nil {
	// 		return ErrBFMissing
	// 	}

	// }
	return nil
}
