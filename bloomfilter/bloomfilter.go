package bloomfilter

import "github.com/DavidHuie/beef/multihash"

type Storage interface {
	GetBit(uint64) (bool, error)
	SetBit(uint64) error
}

type BF struct {
	storage   Storage
	mhash     multihash.Interface
	numHashes uint64
	size      uint64
}

func New(storage Storage, mhash multihash.Interface, numHashes uint64, size uint64) *BF {
	return &BF{
		storage,
		mhash,
		numHashes,
		size,
	}
}

func (b *BF) Insert(v []byte) error {
	hashes := multihash.Hash(b.mhash, v, b.numHashes)

	for _, hash := range hashes {
		index := hash % b.size
		if err := b.storage.SetBit(index); err != nil {
			return err
		}
	}

	return nil
}

func (b *BF) Check(v []byte) (bool, error) {
	hashes := multihash.Hash(b.mhash, v, b.numHashes)

	for _, hash := range hashes {
		index := hash % b.size
		bit, err := b.storage.GetBit(index)
		if err != nil {
			return false, err
		}
		if !bit {
			return false, nil
		}
	}

	return true, nil
}
