package metrics

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// PingExample godoc
// @Summary List Metrics
// @Schemes
// @Description List all users
// @Tags Metrics
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /metrics [get]
func Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"users": "hello world"})
}
