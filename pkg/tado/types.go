package tado

import (
	"time"
)

const (
	ZoneTypeHeating = "HEATING"
)

type GetTokensParams struct {
	Username     string
	Password     string
	ClientSecret string
}

type Tokens struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
	Scope        string `json:"scope"`
	Jti          string `json:"jti"`
}

type Owner struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Homes    []Home `json:"homes"`
}

type Home struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// TimestampedValue represents a value with an associated timestamp
type TimestampedValue struct {
	Value     interface{} `json:"value"`
	Timestamp time.Time   `json:"timestamp"`
}

// OpenWindowDetection represents window detection settings
type OpenWindowDetection struct {
	Supported        bool `json:"supported"`
	Enabled          bool `json:"enabled"`
	TimeoutInSeconds int  `json:"timeoutInSeconds"`
}

// Device represents a physical device in a zone
type Device struct {
	DeviceType       string            `json:"deviceType"`
	SerialNo         string            `json:"serialNo"`
	ShortSerialNo    string            `json:"shortSerialNo"`
	CurrentFwVersion string            `json:"currentFwVersion"`
	ConnectionState  TimestampedValue  `json:"connectionState"`
	MountingState    *TimestampedValue `json:"mountingState,omitempty"`
}

// ZoneStates represents the state mapping for all zones
type ZoneStates struct {
	ZoneStates map[string]ZoneState `json:"zoneStates"`
}

// Zone represents a room or area with connected devices
type Zone struct {
	ID                  int                 `json:"id"`
	Name                string              `json:"name"`
	Type                string              `json:"type"`
	DateCreated         time.Time           `json:"dateCreated"`
	DeviceTypes         []string            `json:"deviceTypes"`
	Devices             []Device            `json:"devices"`
	OpenWindowDetection OpenWindowDetection `json:"openWindowDetection"`
}

// Temperature represents a temperature in multiple units
type Temperature struct {
	Celsius    float64 `json:"celsius"`
	Fahrenheit float64 `json:"fahrenheit"`
}

// TemperatureDataPoint represents a temperature measurement with metadata
type TemperatureDataPoint struct {
	Celsius    float64     `json:"celsius"`
	Fahrenheit float64     `json:"fahrenheit"`
	Timestamp  time.Time   `json:"timestamp"`
	Type       string      `json:"type"`
	Precision  Temperature `json:"precision"`
}

// PercentageDataPoint represents a percentage measurement with metadata
type PercentageDataPoint struct {
	Type       string    `json:"type"`
	Percentage float64   `json:"percentage"`
	Timestamp  time.Time `json:"timestamp"`
}

// SensorPoints represents various sensor measurements
type SensorPoints struct {
	InsideTemperature TemperatureDataPoint `json:"insideTemperature"`
	Humidity          PercentageDataPoint  `json:"humidity"`
}

// ZoneState represents the current state of a zone
type ZoneState struct {
	SensorDataPoints SensorPoints `json:"sensorDataPoints"`
}
