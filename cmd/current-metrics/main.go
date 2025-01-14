package main

import (
	"log"
	"net/http"
	"os"
	"time"

	tado "github.com/attadanta/tado-connect/pkg/tado"
	godotevn "github.com/joho/godotenv"
)

// Print the current state of the zones at home.
//
// Output Fields:
//
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

	log.Printf("Username: %s\n", username)
	log.Printf("Password: %s\n", password)
	log.Printf("Client Secret: %s\n", clientSecret)

	httpClient := &http.Client{
		Timeout: 10 * time.Second,
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

	home := owner.Homes[0]
	log.Printf("Home: %+v\n", home)

	zones, err := tadoClient.GetZones(home.ID)
	if err != nil {
		log.Fatalf("Error getting zones: %s", err)
	}
	log.Printf("Zones: %+v\n", zones)

	states, err := tadoClient.GetZoneStates(home.ID)
	if err != nil {
		log.Fatalf("Error getting zone states: %s", err)
	}
	for zoneId, state := range states.ZoneStates {
		log.Printf("Zone %s: %+v\n", zoneId, state)
	}

	log.Printf("Zone States: %+v\n", states)

}
