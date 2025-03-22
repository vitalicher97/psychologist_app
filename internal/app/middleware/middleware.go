package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
)

// ErrorHandler is a middleware to handle errors centrally.
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next() // Process the request

		if len(c.Errors) > 0 {
			for _, err := range c.Errors {
				if err.Err == pg.ErrNoRows {
					c.JSON(http.StatusNotFound, gin.H{"error": "Resource not found"})
				} else {
					c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				}
				return
			}
		}
	}
}
