package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// swagger: route DELETE /product/{id} products deleteProduct
// Update a products details
//
// responses:
//  201: noContentResponse
//  404: errorResponse
//  501:errorResponse

// Delete handles DELETE requests and removes items from the database
func (p *Products) Delete(c *gin.Context) {
	id := getProductID(c)

	p.l.Println("[DEBUG] deleting record id", id)

	err := data.DeleteProduct(id)
	if err == data.ErrProductNotFound {
		p.l.Println("[ERROR] deleting record id does not exist")

		c.JSON(http.StatusNotFound, &GenericError{Message: err.Error()})
		return
	}

	if err != nil {
		p.l.Println("[ERROR] deleting record", err)

		c.JSON(http.StatusInternalServerError, &GenericError{Message: err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
