package beef

import "github.com/DavidHuie/beef/Godeps/_workspace/src/github.com/boltdb/bolt"

func createBf(tx *bolt.Tx, name string, hashFunction, numHashFunctions, size uint64) error {
	// Check if bucket already exists
	b := tx.Bucket([]byte(name))
	if b != nil {
		return ErrBFAlreadyExists
	}

	// Create bucket
	b, err := tx.CreateBucket(bucketNameForBF(name))
	if err != nil {
		return err
	}

	// Insert blank metadata
	metadata := NewMetadata(hashFunction, numHashFunctions, size)
	serializedMetadata, err := metadata.serialize()
	if err != nil {
		return err
	}
	if err := b.Put(metadataDbKey, serializedMetadata); err != nil {
		return err
	}

	// Initialize pages (we might be able to do this lazily)
	numPages := metadata.pagesNeeded()
	for i := uint64(0); i < numPages; i++ {
		if err := setPage(b, i, make([]byte, pageSize)); err != nil {
			return err
		}

	}

	return nil
}

func insertBf(tx *bolt.Tx, name string, value []byte) error {
	b := tx.Bucket(bucketNameForBF(name))
	if b == nil {
		return ErrBFMissing
	}

	metadata, err := getMetadata(b)
	if err != nil {
		return err
	}

	hashes := metadata.hash.Hash(value, metadata.size)
	for _, hash := range hashes {
		index := hash % metadata.size
		pageIndex, byteIndex, bitIndex := pageMetadata(index)

		page, err := getPage(b, pageIndex)
		if err != nil {
			return err
		}

		byte := page[byteIndex]
		byte = byte | 1<<bitIndex
		page[byteIndex] = byte

		if err := setPage(b, pageIndex, page); err != nil {
			return err
		}
	}

	metadata.InsertCount += 1
	setMetadata(b, metadata)

	return nil
}

func checkBf(tx *bolt.Tx, name string, value []byte) (bool, error) {
	b := tx.Bucket(bucketNameForBF(name))
	if b == nil {
		return false, ErrBFMissing
	}

	metadata, err := getMetadata(b)
	if err != nil {
		return false, err
	}

	hashes := metadata.hash.Hash(value, metadata.size)
	for _, hash := range hashes {
		index := hash % metadata.size
		pageIndex, byteIndex, bitIndex := pageMetadata(index)

		page, err := getPage(b, pageIndex)
		if err != nil {
			return false, err
		}

		byte := (page[byteIndex] >> bitIndex) & 1

		if byte != 1 {
			return false, nil
		}
	}

	return true, nil
}

func destroyBf(tx *bolt.Tx, name string) error {
	b := tx.Bucket(bucketNameForBF(name))
	if b == nil {
		return ErrBFMissing
	}

	return tx.DeleteBucket(bucketNameForBF(name))
}
