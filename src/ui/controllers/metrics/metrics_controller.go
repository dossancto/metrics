package metrics

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1/metrics

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /metrics [get]
func Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"users": "hello world"})
}
