package structure

import (
	"io"
)


type hasher interface {
	// Write (via the embedded io.Writer interface) adds more data to the running hash.
	// It never returns an error.
	io.Writer

	// Sum appends the current hash to b and returns the resulting slice.
	// It does not change the underlying hash state.
	Sum(b []byte) []byte

	// Reset resets the Hash to its initial state.
	Reset()

	// Size returns the number of bytes Sum will return.
	Size() int

	// BlockSize returns the hash's underlying block size.
	// The Write method must be able to accept any amount
	// of data, but it may operate more efficiently if all writes
	// are a multiple of the block size.
	BlockSize() int
}

type entry struct {
	key string
	value interface{}
}

type bucket struct {
	slots [8]*entry
	count int
}

type HashTable struct {
	hasher hasher

	bucketCount int

	buckets []*bucket

	overFlow *HashTable
}

func NewHashTable(hasher hasher, bucketCount int) *HashTable {
	return &HashTable{
		hasher: hasher,
		bucketCount: bucketCount,
		buckets: make([]*bucket, bucketCount),
		overFlow: nil,
	}
}

func (h *HashTable) hash(key string) (int, error) {
	h.hasher.Reset()
	defer h.hasher.Reset()

	_, err := h.hasher.Write([]byte(key))
	if err != nil {
		return 0, err
	}

	hashBytes := h.hasher.Sum(nil)

	hashValue := 0

	for _, b := range hashBytes {
		hashValue = (hashValue << 8) + int(b)
	}

	return hashValue, nil
}

func (h *HashTable) getLOB(data int) int {
	mask := (1 << h.bucketCount) - 1
	return data & mask
}

func (h *HashTable) Insert(key string, value interface{}) error {
	hashedKey, err := h.hash(key)
	if err != nil {
		return err
	}
	lob := h.getLOB(hashedKey)

	nominatedBucket := h.buckets[lob]

	if nominatedBucket == nil {
		slots := [8]*entry{}
		nominatedBucket = &bucket{
			slots: slots,
			count: 0,
		}
		h.buckets[lob] = nominatedBucket
	}
	for i := 0; i < len(nominatedBucket.slots); i++ {
		if nominatedBucket.slots[i] == nil {
			// Empty slot found, insert the new entry
			nominatedBucket.slots[i] = &entry{key: key, value: value}
			nominatedBucket.count++
			return nil
		} else if nominatedBucket.slots[i].key == key {
			// Key already exists, update the value
			nominatedBucket.slots[i].value = value
			return nil
		}
	}

	if h.overFlow == nil {
		h.overFlow = &HashTable{
			hasher: h.hasher,
			bucketCount: h.bucketCount,
			buckets: make([]*bucket, h.bucketCount),
		}
	}

	return h.overFlow.Insert(key, value)
}


func (h *HashTable) Get(key string) (interface{}, bool) {
	index, err := h.hash(key)
	if err != nil {
		return nil, false
	}

	nominatedBucket := h.buckets[index]

	for _, slot := range nominatedBucket.slots {
		if slot != nil && slot.key == key {
			return slot.value, true
		}
	}

	if h.overFlow != nil {
		return h.overFlow.Get(key)
	}
	return nil, false
}


