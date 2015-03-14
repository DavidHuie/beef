package multihash

import (
	"hash/fnv"
	"testing"
)

func TestHash(t *testing.T) {
	hash := fnv.New64()
	hashes := Hash(hash, []byte{1}, 10)

	if len(hashes) != 10 {
		t.Fatalf("invalid number of hashes returned: %v", len(hashes))
	}

	set := make(map[uint64]struct{})
	for _, h := range hashes {
		set[h] = struct{}{}
	}
	if len(set) != 10 {
		// This is not always true, but this check here should catch
		// issues
		t.Fatalf("hashes should be unique")
	}
}
