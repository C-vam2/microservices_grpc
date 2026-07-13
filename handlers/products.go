package handlers

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/microservices_grpc/data"
)

const KeyProduct = "product"

// Products handler for getting and updating products
type Products struct {
	l *log.Logger
	v *data.Validation
}

// NewProduc ts returns a new products handler with the given logger
func NewProducts(l *log.Logger, v *data.Validation) *Products {
	return &Products{l, v}
}

// ErrInvalidProductPath is an error message when the product path is not valid
var ErrInvalidProductPath = fmt.Errorf("Invalid Path, path should be /peoducts/[id]")

// GenericError is a generic error message returned by a server
type GenericError struct {
	Message string `json:"message"`
}

// ValidationError is a collection of validation error messages
type ValidationError struct {
	Messages []string `json:"messages"`
}

// getProductID returns the product ID from the URL
// Panics if cannot convert the id into an integer
// this should never happen as the router ensures that
// this is a valid number

func getProductID(c *gin.Context) int {

	// parse the product id from the url
	value := c.Param("id")

	// convert the id into an integer and return
	id, err := strconv.Atoi(value)

	if err != nil {
		// should never happen
		panic(err)
	}
	return id
}
