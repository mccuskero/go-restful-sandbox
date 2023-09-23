package randomgen

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLowerCaseString(t *testing.T) {
	length := 32
	s := LowerCaseString(length)

	assert.Equal(t, length, len(s))
}
