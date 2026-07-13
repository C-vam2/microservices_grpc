package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/microservices_grpc/data"
)

// MiddlewareValicateProduct validates the product in the request and calls next if ok
func (p *Products) MiddlewareValicateProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		prod := &data.Product{}

		err := c.ShouldBindJSON(prod)
		if err != nil {
			p.l.Println("[ERROR] deserializing product", err)

			c.JSON(http.StatusBadRequest, &GenericError{Message: err.Error()})
			return
		}

		// validate the product
		errs := p.v.Validate(prod)
		if len(errs) != 0 {
			p.l.Println("[ERROR] validating product", err)

			// return the validation messages as an array
			c.JSON(http.StatusUnprocessableEntity, &ValidationError{Messages: errs.Errors()})
			return
		}

		// add the product to the context
		c.Set(KeyProduct, prod)
		c.Next()
	}

}
