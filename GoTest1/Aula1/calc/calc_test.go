package calc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSub(t *testing.T) {
	assert.Equal(t, Sub(7, 5), 2, "A")
	assert.Equal(t, Sub(10, 3), 9, "A")
}
