package theories

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArraySlicing(t *testing.T) {
	assert := assert.New(t)
	origin := []string{"foo", "bar"}
	copied := origin[:]
	copied[0] = "bar"
	fmt.Println(origin, copied)
	assert.NotEqual(copied[0], origin[0])
}
