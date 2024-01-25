package ui

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lu-css/metrics/docs"
	"github.com/lu-css/metrics/src/application/features/metrics"
	"github.com/lu-css/metrics/src/application/middlewares"
	"github.com/lu-css/metrics/src/config"
	"github.com/mvrilo/go-redoc"
	ginredoc "github.com/mvrilo/go-redoc/gin"
	"github.com/newrelic/go-agent/v3/integrations/nrgin"
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
  useThirdParties(r)

	docs.SwaggerInfo.BasePath = "/api/v1"

	openAPIConfig(r)
	r.Use(middlewares.ErrorHandler)

	v1 := r.Group("/api/v1")
	{
		metrics.Routes(v1)

		r.GET("/ping", func(c *gin.Context) { c.JSON(http.StatusOK, "pong") })
	}

	r.Run()
}

func openAPIConfig(r *gin.Engine) {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	redocConfig(r)
}

func redocConfig(r *gin.Engine) {
	r.Use(ginredoc.New(doc))
}

func useThirdParties(r *gin.Engine) {
	r.Use(nrgin.Middleware(config.GetNewRelic()))
}
