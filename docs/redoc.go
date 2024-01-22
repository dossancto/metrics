package docs

import "github.com/mvrilo/go-redoc"

var Doc = redoc.Redoc{
	Title:       "Example API",
	Description: "Example API Description",
	SpecFile:    "./docs/swagger.json", // "./openapi.yaml"
	SpecPath:    "./swagger/doc.json",  // "/openapi.yaml"
	DocsPath:    "/docs",
}
