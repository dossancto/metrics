package features

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lu-css/metrics/docs"
	"github.com/lu-css/metrics/src/application/features/metrics"
	"github.com/mvrilo/go-redoc"
	ginredoc "github.com/mvrilo/go-redoc/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var doc = redoc.Redoc{
	Title:       "Example API",
	Description: "Example API Description",
	SpecFile:    "./docs/swagger.json", // "./openapi.yaml"
	SpecPath:    "./swagger/doc.json",  // "/openapi.yaml"
	DocsPath:    "/docs",
}

func Routes() {
	r := gin.Default()

	docs.SwaggerInfo.BasePath = "/api/v1"

	v1 := r.Group("/api/v1")
	{
		metrics.Routes(v1)

		r.GET("/ping", func(c *gin.Context) { c.JSON(http.StatusOK, "pong") })
	}

	openAPIConfig(r)

	r.Run()
}

func openAPIConfig(r *gin.Engine) {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	redocConfig(r)
}

func redocConfig(r *gin.Engine) {
	r.Use(ginredoc.New(doc))
}
