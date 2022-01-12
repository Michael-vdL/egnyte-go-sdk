package client

import (
	"testing"
)

func Test_ConfigFromFile(t *testing.T) {
	SDKConfig := SDKConfig{
		SubDomain:  "demo",
		APIVersion: "pubapi/v2",
		APIKey:     "0000000011111111",
		APISecret:  "abscdefghijklmnopqrstuvwxyz",
		Username:   "abscdefg",
		Password:   "abscdefg",
	}

	config, err := ConfigurationFromFile("../demo.config.json")
	if err != nil {
		t.Errorf("ConfigurationFromFile failed: %s", err)
	}
	if config.SubDomain != SDKConfig.SubDomain {
		t.Errorf("SubDomain should be %s", SDKConfig.SubDomain)
	}
}

func Test_NewClient(t *testing.T) {
	c := NewClient("subdomain", "apiVersion")
	if c.BaseUrl != "https://subdomain.egnyte.com" {
		t.Errorf("BaseUrl should be https://subdomain.egnyte.com")
	}
}

func Test_Authentication(t *testing.T) {
	config, _ := ConfigurationFromFile("../config.json")
	c := NewClient(config.SubDomain, config.APIVersion)

	err := c.Authenticate(config.Username, config.Password, config.APIKey, config.APISecret)
	if err != nil {
		t.Errorf("Authentication failed: %s", err)
	}
	if c.authToken == "" {
		t.Errorf("authToken should be should be set")
	}
	t.Log(c.authToken)
}
