package ordering

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrdering(t *testing.T) {
	lista := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	listaSorted := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	assert.Equal(t, Sort(lista), listaSorted, "A")
}
