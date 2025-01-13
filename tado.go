package tado

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
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

func jsonResponse(r *http.Response, d any) error {
	if r.StatusCode != 200 {
		return fmt.Errorf("Bad HTTP status %d", r.StatusCode)
	}

	decoder := json.NewDecoder(r.Body)
	return decoder.Decode(&d)
}

func GetBearerToken(c *http.Client, p GetTokensParams) (Tokens, error) {
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
