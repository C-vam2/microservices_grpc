package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Wrap function takes http handler and return a gin handler
func Wrap(h http.Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
