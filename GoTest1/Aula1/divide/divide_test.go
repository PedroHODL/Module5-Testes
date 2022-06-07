package divide

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDivide(t *testing.T) {
	res, _ := Division(10, 2)
	assert.Equal(t, res, 5, "A divisão falhou")
}

func TestDivideFail(t *testing.T) {
	_, err := Division(50, 0)
	assert.NotNil(t, err, "Deveria ocorrer um erro ao usar 0")
}
