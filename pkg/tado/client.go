package tado

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func jsonResponse(r *http.Response, d any) error {
	if r.StatusCode != 200 {
		return fmt.Errorf("Bad HTTP status %d", r.StatusCode)
	}

	decoder := json.NewDecoder(r.Body)
	return decoder.Decode(&d)
}

type TadoClient struct {
	client       *http.Client
	auth         *Tokens
	clientSecret string
}

func NewTadoClient(c *http.Client, clientSecret string, t Tokens) *TadoClient {
	return &TadoClient{
		client:       c,
		auth:         &t,
		clientSecret: clientSecret,
	}
}

// Authenticate obtains an access token from the `oauth/token` resource.
func Authenticate(c *http.Client, p GetTokensParams) (Tokens, error) {
	f := url.Values{}
	f.Add("client_id", "tado-web-app")
	f.Add("grant_type", "password")
	f.Add("scope", "home.user")
	f.Add("username", p.Username)
	f.Add("password", p.Password)
	f.Add("client_secret", p.ClientSecret)

	req, err := http.NewRequest("POST", "https://auth.tado.com/oauth/token", strings.NewReader(f.Encode()))
	if err != nil {
		return Tokens{}, err
	}
	req.Header.Add("content-type", "application/x-www-form-urlencoded")

	res, err := c.Do(req)
	if err != nil {
		return Tokens{}, err
	}
	defer res.Body.Close()

	tokens := Tokens{}
	err = jsonResponse(res, &tokens)
	if err != nil {
		return Tokens{}, err
	}

	return tokens, nil
}

func (t *TadoClient) GetMe() (Owner, error) {
	req, err := http.NewRequest("GET", "https://my.tado.com/api/v2/me", nil)
	if err != nil {
		return Owner{}, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", t.auth.AccessToken))

	res, err := t.client.Do(req)
	if err != nil {
		return Owner{}, err
	}
	defer res.Body.Close()

	owner := Owner{}
	err = jsonResponse(res, &owner)
	if err != nil {
		return Owner{}, err
	}

	return owner, nil
}

func GetRefreshToken(c *http.Client, p GetRefreshTokenParams) (Tokens, error) {
	f := url.Values{}
	f.Add("client_id", "tado-web-app")
	f.Add("grant_type", "refresh_token")
	f.Add("client_secret", p.ClientSecret)
	f.Add("refresh_token", p.RefreshToken)

	req, err := http.NewRequest("POST", "https://auth.tado.com/oauth/token", strings.NewReader(f.Encode()))
	if err != nil {
		return Tokens{}, err
	}
	req.Header.Add("content-type", "application/x-www-form-urlencoded")

	res, err := c.Do(req)
	if err != nil {
		return Tokens{}, err
	}
	defer res.Body.Close()

	tokens := Tokens{}
	err = jsonResponse(res, &tokens)
	if err != nil {
		return Tokens{}, err
	}

	return tokens, nil
}

func (t *TadoClient) GetZones(homeID int) ([]Zone, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://my.tado.com/api/v2/homes/%d/zones", homeID), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", t.auth.AccessToken))

	res, err := t.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	zones := []Zone{}
	err = jsonResponse(res, &zones)
	if err != nil {
		return nil, err
	}

	return zones, nil
}

func (t *TadoClient) GetZoneState(homeID int, zoneId string) (ZoneState, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://my.tado.com/api/v2/homes/%d/zones/%s/state", homeID, zoneId), nil)
	if err != nil {
		return ZoneState{}, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", t.auth.AccessToken))
	res, err := t.client.Do(req)
	if err != nil {
		return ZoneState{}, err
	}
	defer res.Body.Close()

	zone := ZoneState{}
	err = jsonResponse(res, &zone)
	if err != nil {
		return ZoneState{}, err
	}

	return zone, nil
}

func (t *TadoClient) GetZoneStates(homeID int) (ZoneStates, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://my.tado.com/api/v2/homes/%d/zoneStates", homeID), nil)
	if err != nil {
		return ZoneStates{}, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", t.auth.AccessToken))

	res, err := t.client.Do(req)
	if err != nil {
		return ZoneStates{}, err
	}
	defer res.Body.Close()

	zones := ZoneStates{}
	err = jsonResponse(res, &zones)
	if err != nil {
		return ZoneStates{}, err
	}

	return zones, nil
}
