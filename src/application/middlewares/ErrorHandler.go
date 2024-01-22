package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lu-css/metrics/src/application/domain/exceptions"
)

func ErrorHandler(c *gin.Context) {
	c.Next()

	for _, err := range c.Errors {
		switch err.Err.(type) {

		case *exceptions.ErrNotFound:
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		}

	}

	c.JSON(http.StatusInternalServerError, "")
}
