package beef

import (
	"encoding/binary"
	"fmt"
)

func bucketNameForBF(name string) []byte {
	return []byte(fmt.Sprintf("bf:%v", name))
}

func pageKey(n uint64) []byte {
	prefix := []byte("p:")
	nBytes := make([]byte, 8)
	binary.PutUvarint(nBytes, n)
	return append(prefix, nBytes...)
}
