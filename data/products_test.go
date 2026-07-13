package data

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductMissingNameReturnsErr(t *testing.T) {
	p := Product{
		Price: 1.14,
	}
	v := NewValidation()
	err := v.Validate(p)
	assert.Len(t, err, 1)
}

func TestProductMissingPriceReturrnsErr(t *testing.T) {
	p := Product{
		Name:  "abc",
		Price: -1,
	}

	v := NewValidation()
	err := v.Validate(p)
	assert.Len(t, err, 1)
}

func TestProductInvalidSKUReturnsErr(t *testing.T) {
	p := Product{
		Name:  "abc",
		Price: 1.22,
		SKU:   "abc",
	}

	v := NewValidation()
	err := v.Validate(p)
	assert.Len(t, err, 1)
}

func TestValidProductDoesNOTReturnsErr(t *testing.T) {
	p := Product{
		Name:  "abc",
		Price: 1.22,
		SKU:   "abc-def-ghi",
	}

	v := NewValidation()
	err := v.Validate(p)
	assert.Len(t, err, 1)
}
