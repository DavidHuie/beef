package multihash

import (
	"encoding/binary"
	"hash"
)

func hashPadding(n uint64) []byte {
	buf := make([]byte, 4)
	binary.PutUvarint(buf, n)
	return buf
}

func payloadsWithPadding(b []byte, n uint64) [][]byte {
	payloads := make([][]byte, 0)
	size := len(b)

	for i := uint64(0); i < n; i++ {
		c := make([]byte, size)
		copy(c, b)
		c = append(c, hashPadding(i)...)
		payloads = append(payloads, c)
	}

	return payloads
}

func Hash(h hash.Hash64, value []byte, numHashes uint64) []uint64 {
	hashes := make([]uint64, 0)
	payloads := payloadsWithPadding(value, numHashes)

	for _, payload := range payloads {
		h.Reset()
		h.Write(payload)
		hash := h.Sum64()

		hashes = append(hashes, hash)
	}

	return hashes
}
