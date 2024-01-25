package config

import (
	"github.com/mvrilo/go-redoc"
	"github.com/swaggo/swag"
)

func GetSwaggerConfig(s swag.Spec) *swag.Spec {
	return &swag.Spec{
		Version:          "1.0",
		Host:             "localhost",
		BasePath:         "/swagger",
		Schemes:          []string{},
		Title:            "Metrics Swagger",
		Description:      "Application to check metrics",
		InfoInstanceName: "swagger",
		SwaggerTemplate:  s.SwaggerTemplate,
		LeftDelim:        "{{",
		RightDelim:       "}}",
	}
}

func GetRedocConfig() redoc.Redoc {
	return redoc.Redoc{
		Title:       "Example API",
		Description: "Example API Description",
		SpecFile:    "./docs/swagger.json", // "./openapi.yaml"
		SpecPath:    "./swagger/doc.json",  // "/openapi.yaml"
		DocsPath:    "/docs",
	}
}
