package config

import (
	"log"
	"os"

	"github.com/newrelic/go-agent/v3/newrelic"
)

func GetNewRelic() *newrelic.Application {
	appName := os.Getenv("NEW_RELIC_APP_NAME")
	license := os.Getenv("NEW_RELIC_LICENSE")

	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName(appName),
		newrelic.ConfigLicense(license),
		newrelic.ConfigAppLogForwardingEnabled(true),
	)

	if err != nil {
		log.Println("WARN: Error while configuring new relic")
		log.Println(err)
	}

	return app
}
