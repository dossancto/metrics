package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/lu-css/metrics/src/application/ui"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("unable to load .env file: %e", err)
	}
	ui.Routes()
}
