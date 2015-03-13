package beef

import (
	"encoding/json"
	"hash/fnv"

	"github.com/DavidHuie/beef/Godeps/_workspace/src/github.com/boltdb/bolt"
	"github.com/DavidHuie/beef/multihash"
)

type Metadata struct {
	InsertCount      uint64
	hash             *multihash.MultiHash
	hashFunction     uint64
	numHashFunctions uint64
	size             uint64
}

func NewMetadata(hashFunction, numHashFunctions, size uint64) *Metadata {
	return &Metadata{
		InsertCount:      0,
		hashFunction:     hashFunction,
		numHashFunctions: numHashFunctions,
		size:             size,
	}
}

func getMetadata(bucket *bolt.Bucket) (*Metadata, error) {
	var metadataBytes []byte

	metadataBytes = bucket.Get(metadataDbKey)
	if metadataBytes == nil {
		return nil, ErrMetadataMissing
	}

	var metadata Metadata
	if err := json.Unmarshal(metadataBytes, &metadata); err != nil {
		return nil, err
	}

	metadata.initialize()

	return &metadata, nil
}

func setMetadata(bucket *bolt.Bucket, metadata *Metadata) error {
	serialized, err := metadata.serialize()
	if err != nil {
		return err
	}
	return bucket.Put(metadataDbKey, serialized)
}

func (m *Metadata) initialize() {
	m.hash = multihash.New(fnv.New64())
}

func (m *Metadata) serialize() ([]byte, error) {
	bytes, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func (m *Metadata) pagesNeeded() uint64 {
	return (m.size % pageSize) + 1
}
