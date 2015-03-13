package multihash

import (
	"encoding/binary"
	"hash"
)

// Uses a hash function to apply generate n hashes
// for a given payload
type MultiHash struct {
	hash hash.Hash64
}

func New(hash hash.Hash64) *MultiHash {
	return &MultiHash{
		hash,
	}
}

func hashPadding(n int) []byte {
	buf := make([]byte, 4)
	binary.PutVarint(buf, int64(n))
	return buf
}

func payloadsWithPadding(b []byte, n int) [][]byte {
	payloads := make([][]byte, 0)
	size := len(b)

	for i := 0; i < n; i++ {
		c := make([]byte, size)
		copy(c, b)
		c = append(c, hashPadding(i)...)
		payloads = append(payloads, c)
	}

	return payloads
}

func (m *MultiHash) hashSingle(b []byte) uint64 {
	m.hash.Reset()
	m.hash.Write(b)

	return m.hash.Sum64()
}

func (m *MultiHash) Hash(b []byte, n int) []uint64 {
	hashes := make([]uint64, 0)
	payloads := payloadsWithPadding(b, n)

	for _, payload := range payloads {
		hashes = append(hashes, m.hashSingle(payload))
	}

	return hashes
}
