package server

import (
	"encoding/json"
	"errors"
	"hash"
	"hash/fnv"

	"github.com/DavidHuie/beef/Godeps/_workspace/src/github.com/boltdb/bolt"
)

var (
	metadataDbKey = []byte("m")

	ErrMetadataMissing = errors.New("metadata is missing")
)

type Metadata struct {
	InsertCount      uint64
	NumHashFunctions uint64
	Size             uint64

	hash hash.Hash64
}

func (m *Metadata) initialize() {
	m.hash = fnv.New64()
}

func NewMetadata(numHashFunctions, size uint64) *Metadata {
	m := &Metadata{
		InsertCount:      0,
		NumHashFunctions: numHashFunctions,
		Size:             size,
	}
	m.initialize()
	return m
}

func (s *Server) GetMetadata(bucket *bolt.Bucket) (*Metadata, error) {
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

func (s *Server) SetMetadata(bucket *bolt.Bucket, metadata *Metadata) error {
	serialized, err := metadata.serialize()
	if err != nil {
		return err
	}
	return bucket.Put(metadataDbKey, serialized)
}

func (m *Metadata) serialize() ([]byte, error) {
	bytes, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}
