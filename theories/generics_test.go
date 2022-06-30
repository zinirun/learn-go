package theories

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumInts(t *testing.T) {
	assert := assert.New(t)
	ints := map[string]int{"first": 1, "second": 2}
	s := Sum(ints)
	assert.Equal(s, 3)
}

func TestSumFloats(t *testing.T) {
	assert := assert.New(t)
	floats := map[string]float64{"first": 1.0, "second": 2.0}
	s := Sum(floats)
	assert.Equal(s, 3.0)
}
