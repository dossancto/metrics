package ui

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lu-css/metrics/docs"
	"github.com/lu-css/metrics/src/application/features/metrics"
	"github.com/lu-css/metrics/src/application/middlewares"
	"github.com/lu-css/metrics/src/config"

	// "github.com/lu-css/metrics/src/config"
	ginredoc "github.com/mvrilo/go-redoc/gin"

	// "github.com/newrelic/go-agent/v3/integrations/nrgin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Routes() {
	r := gin.Default()
	useThirdParties(r)

	docs.SwaggerInfo = config.GetSwaggerConfig(*docs.SwaggerInfo)

	openAPIConfig(r)
	r.Use(middlewares.ErrorHandler)

	metrics.Routes(&r.RouterGroup)

	r.Group("/api/v1")
	{
		r.GET("/ping", func(c *gin.Context) { c.JSON(http.StatusOK, "pong") })
	}

	r.Run()
}

func openAPIConfig(r *gin.Engine) {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	redocConfig(r)
}

func redocConfig(r *gin.Engine) {
	r.Use(ginredoc.New(config.GetRedocConfig()))
}

func useThirdParties(r *gin.Engine) {
	// r.Use(nrgin.Middleware(config.GetNewRelic()))
}
