package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/microservices_grpc/data"
)

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

func (p *Products) GetProducts(c *gin.Context) {
	lp := data.GetProducts()
	err := lp.ToJSON(c.Writer)

	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Unable to parse the product JSON",
		})
		return
	}
}

func (p *Products) AddProduct(c *gin.Context) {
	p.l.Println("POST request triggered!!")

	prod := &data.Product{}
	if err := c.ShouldBindJSON(&prod); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	data.AddProduct(prod)
}

func (p *Products) UpdateProduct(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		p.l.Println(err)
		p.l.Println(c.Param("id"))
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid productID",
		})
		return
	}

	value, exists := c.Get("product")

	if !exists {
		c.JSON(http.StatusInternalServerError, "Product not found in the request")
		return
	}

	prod := value.(*data.Product)
	err = data.UpdateProduct(id, prod)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

}

func (p *Products) MiddlewareProductValidation(c *gin.Context) {
	prod := &data.Product{}
	if err := c.ShouldBindJSON(&prod); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	err := prod.Validate()
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]any{
			"error": err.Error(),
		})
	}
	c.Set("product", prod)
	c.Next()
}
