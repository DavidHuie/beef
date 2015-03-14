package server

import (
	"errors"
	"fmt"

	"github.com/DavidHuie/beef/Godeps/_workspace/src/github.com/boltdb/bolt"
	"github.com/DavidHuie/beef/bloomfilter"
	"github.com/DavidHuie/beef/storage/boltdb"
)

var (
	ErrBFMissing = errors.New("bloom filter not found")
)

func bucketNameForBF(name string) []byte {
	return []byte(fmt.Sprintf("bf:%v", name))
}

func (s *Server) Insert(tx *bolt.Tx, name string, value []byte) error {
	b := tx.Bucket(bucketNameForBF(name))
	if b == nil {
		return ErrBFMissing
	}

	metadata, err := GetMetadata(b)
	if err != nil {
		return err
	}

	storage := boltdb.New(b)
	bf := bloomfilter.New(
		storage,
		metadata.hash,
		metadata.Size,
		metadata.NumHashFunctions,
	)
	if err := bf.Insert(value); err != nil {
		return err
	}

	metadata.InsertCount += 1
	if err := SetMetadata(b, metadata); err != nil {
		return err
	}

	return nil
}

func (s *Server) Check(tx *bolt.Tx, name string, value []byte) (bool, error) {
	b := tx.Bucket(bucketNameForBF(name))
	if b == nil {
		return false, ErrBFMissing
	}

	metadata, err := GetMetadata(b)
	if err != nil {
		return false, err
	}

	storage := boltdb.New(b)
	bf := bloomfilter.New(
		storage,
		metadata.hash,
		metadata.Size,
		metadata.NumHashFunctions,
	)
	check, err := bf.Check(value)
	if err != nil {
		return false, err
	}

	return check, err
}
