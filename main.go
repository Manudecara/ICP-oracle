package main

import (
	"time"

	framework "github.com/hyplabs/dfinity-oracle-framework"
	"github.com/hyplabs/dfinity-oracle-framework/models"
)

func main() {
	tokyoEndpoints := []models.Endpoint{
		{
			Endpoint: "http://api.weatherapi.com/v1/current.json?key=fa0e65846eba499eaa4104508220201&q=Tokyo,JP",
			JSONPaths: map[string]string{
				"temperature_celsius": "$.current.temp_c",
			},
		},
		{
			Endpoint: "https://api.weatherbit.io/v2.0/current?key=fa0e65846eba499eaa4104508220201&city=Tokyo&country=JP",
			JSONPaths: map[string]string{
				"temperature_celsius": "$.data[0].temp",
			},
		},
	}
	delhiEndpoints := []models.Endpoint{
		{
			Endpoint: "http://api.weatherapi.com/v1/current.json?key=fa0e65846eba499eaa4104508220201&q=Delhi,IN",
			JSONPaths: map[string]string{
				"temperature_celsius": "$.current.temp_c",
			},
		},
		{
			Endpoint: "https://api.weatherbit.io/v2.0/current?key=fa0e65846eba499eaa4104508220201&city=Delhi&country=IN",
			JSONPaths: map[string]string{
				"temperature_celsius": "$.data[0].temp",
			},
		},
	}
	config := models.Config{
		CanisterName:   "sample_oracle",
		UpdateInterval: 5 * time.Minute,
	}
	engine := models.Engine{
		Metadata: []models.MappingMetadata{
			{Key: "Tokyo", Endpoints: tokyoEndpoints},
			{Key: "Delhi", Endpoints: delhiEndpoints},
		},
	}
	oracle := framework.NewOracle(&config, &engine)
	oracle.Bootstrap()
	oracle.Run()
}
