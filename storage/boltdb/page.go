package boltdb

import (
	"encoding/binary"

	"github.com/DavidHuie/beef/Godeps/_workspace/src/github.com/boltdb/bolt"
)

var (
	PageSize = uint64(32) // bytes
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

func getPage(bucket *bolt.Bucket, page uint64) []byte {
	p := bucket.Get(pageKey(page))
	if p == nil {
		return make([]byte, PageSize)
	}
	return p
}
