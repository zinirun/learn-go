package theories

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	assert := assert.New(t)
	ints := map[string]int{"first": 1, "second": 2}
	s := Sum(ints)
	assert.Equal(s, 3)
}
