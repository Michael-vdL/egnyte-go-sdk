package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const (
  DefaultSubDomain = "demo" // Sample Testing Subdomain
  DefaultEgnyteURL = "egnyte.com"
  DefaultEgnyteAPIv1 = "/pubapi/v1/"
  DefaultEgnyteAPIv2 ="/pubapi/v2/"
  DefaultEgnyteAuthAPI = "/pubauth/token"
)

type Client struct {
  BaseUrl string
  APIVersion string
  httpClient *http.Client
  authToken string
}



func (c *Client) Authenticate(username, password, clientKey, clientSecret string) error {
  authURL := fmt.Sprintf("%s%s", c.BaseUrl, DefaultEgnyteAuthAPI)

  authData := url.Values{}
  authData.Set("grant_type", "password")
  authData.Set("username", username)
  authData.Set("password", password)
  authData.Set("client_key", clientKey)
  authData.Set("client_secret", clientSecret)


  req, err := http.NewRequest("POST", authURL, strings.NewReader(authData.Encode()))
  if err != nil {
    return err
  }
  req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
  req.Header.Add("Content-Length", strconv.Itoa(len(authData.Encode())))

  res, err := c.httpClient.Do(req)
  if err != nil {
    return err
  }
  defer res.Body.Close()
  if res.StatusCode != 200 {
    return fmt.Errorf("authentication failed. status code: %d", res.StatusCode)
  }

  var authResponse AuthResponse
  err = json.NewDecoder(req.Body).Decode(&authResponse)
  if err != nil {
    return err
  }

  c.authToken = authResponse.AccessToken

  return nil
}


func NewClient(subDomain, apiVersion string) *Client {
  url := fmt.Sprintf("https://%s.%s", subDomain, DefaultEgnyteURL)

  return &Client{
    BaseUrl: url,
    APIVersion: apiVersion,
    httpClient: http.DefaultClient,
    authToken: "",
  }
}