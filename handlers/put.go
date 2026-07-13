package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/microservices_grpc/data"
)

// swagger:route PUT /products products updateProduct
// Update a products details
//
// responses:
//  201: noContentResponse
//  404: errorResponse
//  422: errorValidation

// Update handles PUT requests to update products
func (p *Products) Update(c *gin.Context) {

	// fetch the product from the context
	value, exists := c.Get(KeyProduct)

	if !exists {
		p.l.Println("[ERROR] product not found in gin context")

		c.JSON(http.StatusNotFound, &GenericError{Message: "Unable to update product"})
		return
	}

	prod := value.(*data.Product)
	p.l.Println("[DEBUG] updating record id", prod.ID)

	err := data.UpdateProduct(prod)
	if err == data.ErrProductNotFound {
		p.l.Println("[ERROR] product not found", err)

		c.JSON(http.StatusNotFound, &GenericError{Message: "Product not found in database"})
		return
	}

	// write the no content success header
	c.Status(http.StatusNoContent)
}
