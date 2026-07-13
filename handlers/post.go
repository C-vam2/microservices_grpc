package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/microservices_grpc/data"
)

// swagger:route POST /products products createProduct
// Create a new product
//
// response:
//  200: productResponse
//  422: errorValidation
//  501: errorResponse

// Create handles POST requests to add new products
func (p *Products) Create(c *gin.Context) {
	// fetch the product from the context
	value, exists := c.Get(KeyProduct)

	if !exists {
		p.l.Println("[ERROR] product object not found in the gin context")

		c.JSON(http.StatusNotFound, &GenericError{Message: "Unable to add new product"})
		return
	}

	prod := value.(*data.Product)

	p.l.Println("[DEBUG] Inserting product: %#v\n", value)
	data.AddProduct(prod)
}
