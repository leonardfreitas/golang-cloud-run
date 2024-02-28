package entity_test

import (
	"testing"

	"github.com/leonardfreitas/golang-cloud-run/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestNormalizePostalCode(t *testing.T) {
	normalized, err := entity.NormalizePostalCode("12345-678")
	assert.NoError(t, err)
	assert.Equal(t, "12345678", normalized)

	normalized, err = entity.NormalizePostalCode("")
	assert.EqualError(t, err, "cep vazio")

	normalized, err = entity.NormalizePostalCode("1234")
	assert.EqualError(t, err, "formado do cep inválido")

	normalized, err = entity.NormalizePostalCode("ABC12345678")
	assert.Equal(t, "12345678", normalized)
	assert.NoError(t, err)
}
