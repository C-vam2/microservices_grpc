package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/microservices_grpc/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
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
	c.Set("product", prod)
	c.Next()
}
