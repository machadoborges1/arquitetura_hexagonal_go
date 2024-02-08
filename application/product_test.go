package application_test

import (
	"testing"

	"github.com/machadoborges1/arquitetura_hexagonal_go/application"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func TestProduct_Enable(t *testing.T) {
	product := application.Product{}
	product.Name = "Hello"
	product.Status = application.DISABLED
	product.Price = 10

	err := product.Enable()
	assert.Nil(t, err)

	product.Price = 0
	err = product.Enable()
	assert.Equal(t, "the price must be greater than zero to enable the product", err.Error())
}

func TestProduct_Disable(t *testing.T) {
	product := application.Product{}
	product.Name = "Hello"
	product.Status = application.ENABLED
	product.Price = 0

	err := product.Disable()
	assert.Nil(t, err)

	product.Price = 10
	err = product.Disable()
	assert.Equal(t, "the price must be zero in order to have the product disabled", err.Error())
}

func TestProduct_IsValid(t *testing.T) {
	product := application.Product{}
	product.ID = uuid.NewV4().String()
	product.Name = "hello"
	product.Status = application.DISABLED
	product.Price = 10

	_, err := product.IsValid()
	assert.Nil(t, err)

	product.Status = "INVALID"
	_, err = product.IsValid()
	assert.Equal(t, "the status must be enabled or disabled", err.Error())

	product.Status = application.ENABLED
	_, err = product.IsValid()
	assert.Nil(t, err)

	product.Price = -10
	_, err = product.IsValid()
	assert.Equal(t, "the price must be greater or equal zero", err.Error())
}
