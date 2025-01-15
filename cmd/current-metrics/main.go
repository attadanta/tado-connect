package main

import (
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	tado "github.com/attadanta/tado-connect/pkg/tado"
	godotevn "github.com/joho/godotenv"
)

func fetchZoneStatesAndPrint(c *tado.TadoClient, ticker time.Ticker, done chan bool, homeID int) {
	for {
		select {
		case <-done:
			log.Printf("Stopping zone state fetcher")
			return
		case <-ticker.C:
			log.Printf("Getting new zone states")
			states, err := c.GetZoneStates(homeID)
			if err != nil {
				log.Fatalf("Error getting zone states: %s", err)
			}
			for zoneId, state := range states.ZoneStates {
				log.Printf("Zone %s: %+v\n", zoneId, state)
			}
		}
	}
}

// Print the current state of the zones at home.
//
// Output Fields:
//
// - Timestamp
// - ZoneId
// - Sensor
// - InsideTemperature
// - Humidity
func main() {
	err := godotevn.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	username := os.Getenv("TADO_USERNAME")
	password := os.Getenv("TADO_PASSWORD")
	clientSecret := os.Getenv("TADO_CLIENT_SECRET")

	timeoutRaw := os.Getenv("HTTP_CLIENT_TIMEOUT")
	timeout, err := time.ParseDuration(timeoutRaw)
	if err != nil {
		log.Fatalf("Error parsing HTTP_CLIENT_TIMEOUT: %s", err)
	}

	tickerPeriodRaw := os.Getenv("REFETCH_PERIOD")
	log.Printf("Ticker period: %s\n", tickerPeriodRaw)
	tickerPeriod, err := time.ParseDuration(tickerPeriodRaw)
	if err != nil {
		log.Fatalf("Error parsing REFETCH_PERIOD: %s", err)
	}

	httpClient := &http.Client{
		Timeout: timeout,
	}

	tokens, err := tado.GetBearerToken(httpClient, tado.GetTokensParams{
		Username:     username,
		Password:     password,
		ClientSecret: clientSecret,
	})
	if err != nil {
		log.Fatalf("Error getting bearer token: %s", err)
	}

	tadoClient := tado.NewTadoClient(httpClient, tokens)
	owner, err := tadoClient.GetMe()
	if err != nil {
		log.Fatalf("Error getting owner: %s", err)
	}
	log.Printf("Owner: %+v\n", owner)

	homes := owner.Homes
	if len(homes) == 0 {
		log.Fatalf("No homes found")
	}

	ticker := time.NewTicker(tickerPeriod)
	done := make(chan bool)

	home := owner.Homes[0]
	log.Printf("Home: %+v\n", home)

	fetchZoneStatesAndPrint(tadoClient, *ticker, done, home.ID)

	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()
}
