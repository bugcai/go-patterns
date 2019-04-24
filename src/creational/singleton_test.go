package creational

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSingleton(t *testing.T) {
	s := NewInsance()
	s["this"] = "that"

	s2 := NewInsance()
	assert.Equal(t, s["this"], s2["this"])
}
