package creational

import "testing"

func TestFactory(t *testing.T) {
	s := NewStore(MemoryStorage)
	s.Open("file")
}
