package metrics

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lu-css/metrics/src/application/features/metrics/usecases"
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

// @BasePath /api/v1

// Create
// @Summary Create Metric
// @Schemes
// @Description Creates a new Metric
// @Tags Metrics
// @Accept json
// @Produce json
// @Success 200 {Metric} Metric
// @Router /metrics [POST]
func Create(c *gin.Context) {
	result, err := usecases.Create("salve")

	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, result)
}
