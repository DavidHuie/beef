package beef

import (
	"errors"

	"github.com/DavidHuie/beef/Godeps/_workspace/src/github.com/boltdb/bolt"
)

var (
	ErrPageMissing = errors.New("page does not exist")
)

func pageMetadata(index uint64) (uint64, uint64, uint64) {
	globalByteIndex := index / 8
	pageIndex := globalByteIndex / pageSize
	byteIndex := globalByteIndex % pageSize
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
