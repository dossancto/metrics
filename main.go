package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"

	docs "github.com/lu-css/metrics/docs"

	metrics "github.com/lu-css/metrics/src/ui/controllers/metrics"
	ginredoc "github.com/mvrilo/go-redoc/gin"
)

func main() {
	r := gin.Default()

	docs.SwaggerInfo.BasePath = "/api/v1"

	v1 := r.Group("/api/v1")
	{
		metrics.Routes(v1)

		r.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, "eae pogao")
		})
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.Use(ginredoc.New(docs.Doc))
	r.Run()
}
