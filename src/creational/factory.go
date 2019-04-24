package creational

import "io"

type Store interface {
	Open(string) (io.ReadWriteCloser, error)
}

type StorageType int

const (
	DiskStorage StorageType = 1 << iota
	TempStorage
	MemoryStorage
)

func NewStore(t StorageType) Store {
	switch t {
	case MemoryStorage:
		return newMemoryStorage()
	case DiskStorage:
		return newDiskStorage()
	default:
		return newTempStorage()
	}
}

type tempStore struct {
}

func (c *tempStore) Open(string) (io.ReadWriteCloser, error) {
	return nil, nil
}

func newTempStorage() Store {
	return &tempStore{}
}

type diskStore struct {
}

func (c *diskStore) Open(string) (io.ReadWriteCloser, error) {
	return nil, nil
}

func newDiskStorage() Store {
	return &diskStore{}
}

type memoryStore struct {
}

func (c *memoryStore) Open(string) (io.ReadWriteCloser, error) {
	return nil, nil
}

func newMemoryStorage() Store {
	return &memoryStore{}
}
