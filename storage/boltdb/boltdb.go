package boltdb

import "github.com/DavidHuie/beef/Godeps/_workspace/src/github.com/boltdb/bolt"

type DB struct {
	bucket *bolt.Bucket
}

func New(bucket *bolt.Bucket) *DB {
	return &DB{bucket}
}

func (db *DB) GetBit(n uint64) (bool, error) {
	pageIndex, byteIndex, bitIndex := pageMetadata(n)
	page, err := getPage(db.bucket, pageIndex)
	if err != nil {
		return false, err
	}

	byte := (page[byteIndex] >> bitIndex) & 1

	return byte == 1, nil
}

func (db *DB) SetBit(n uint64) error {
	pageIndex, byteIndex, bitIndex := pageMetadata(n)
	page, err := getPage(db.bucket, pageIndex)
	if err != nil {
		return err
	}

	byte := page[byteIndex]
	byte = byte | 1<<bitIndex
	page[byteIndex] = byte

	if err := db.bucket.Put(pageKey(pageIndex), page); err != nil {
		return err
	}

	return nil
}
