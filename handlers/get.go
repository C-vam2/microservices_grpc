package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/microservices_grpc/data"
)

// swagger:router GET /products products listProducts
// Return a list of products from the database
// responses:
//  200: productsResponse

// ListAll handles GET request and returns all current products
func (p *Products) ListAll(c *gin.Context) {
	p.l.Println("[DEBUG] get all records")

	prods := data.GetProducts()

	err := c.ShouldBindJSON(prods)

	if err != nil {
		// we should never be here but log the error just incase
		p.l.Println("[ERROR] serializing product", err)
	}
}

// swagger:route GET /products/{id} products listSingle
// Return a list of products from the databse
// responses:
//  200: productResponse
//  404: errorResponse

// ListSingle handles GET requests
func (p *Products) ListSingle(c *gin.Context) {
	id := getProductID(c)

	p.l.Println("[DEBUG] get record id", id)

	prod, err := data.GetProductByID(id)

	switch err {
	case nil:

	case data.ErrProductNotFound:
		p.l.Println("[ERROR] fetching product", err)

		c.JSON(http.StatusNotFound,
			&GenericError{Message: err.Error()})
		return
	default:
		p.l.Println("[ERROR] fetching product", err)

		c.JSON(http.StatusInternalServerError, &GenericError{Message: err.Error()})
		return
	}

	err = c.ShouldBindJSON(prod)
	if err != nil {
		// we should never be here but log the error just incase
		p.l.Println("[ERROR] serializing product", err)
	}

}
