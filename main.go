package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	docs "github.com/lu-css/metrics/docs"

	metrics "github.com/lu-css/metrics/src/ui/controllers/metrics"
	"github.com/mvrilo/go-redoc"
	ginredoc "github.com/mvrilo/go-redoc/gin"
)

var doc = redoc.Redoc{
	Title:       "Example API",
	Description: "Example API Description",
	SpecFile:    "./docs/swagger.json", // "./openapi.yaml"
	SpecPath:    "./swagger/doc.json",  // "/openapi.yaml"
	DocsPath:    "/docs",
}

func main() {
	r := gin.Default()

	docs.SwaggerInfo.BasePath = "/api/v1"

	v1 := r.Group("/api/v1")
	{
		metrics.Routes(v1)

		r.GET("/ping", func(c *gin.Context) { c.JSON(http.StatusOK, "eae pogao") })
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.Use(ginredoc.New(doc))
	r.Run()
}
