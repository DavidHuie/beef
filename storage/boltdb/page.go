package boltdb

import (
	"encoding/binary"
	"errors"

	"github.com/boltdb/bolt"
)

var (
	PageSize = uint64(32) // bytes

	ErrPageMissing = errors.New("page does not exist")
)

func SetPageSize(n uint64) {
	PageSize = n
}

func pageKey(n uint64) []byte {
	prefix := []byte("p:")
	nBytes := make([]byte, 8)
	binary.PutUvarint(nBytes, n)
	return append(prefix, nBytes...)
}

func pageMetadata(index uint64) (uint64, uint64, uint64) {
	globalByteIndex := index / 8
	pageIndex := globalByteIndex / PageSize
	byteIndex := globalByteIndex % PageSize
	bitIndex := index % 8

	return pageIndex, byteIndex, bitIndex
}

func getPage(bucket *bolt.Bucket, page uint64) ([]byte, error) {
	p := bucket.Get(pageKey(page))
	if p == nil {
		return nil, ErrPageMissing
	}
	return p, nil
}

func setPage(bucket *bolt.Bucket, page uint64, value []byte) error {
	return bucket.Put(pageKey(page), value)
}
