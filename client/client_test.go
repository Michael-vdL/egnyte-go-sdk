package client

import (
	"testing"
)

func Test_NewClient(t *testing.T) {
  c := NewClient("subdomain", "apiVersion")
  if c.BaseUrl != "https://subdomain.egnyte.com" {
    t.Errorf("BaseUrl should be https://subdomain.egnyte.com")
  }
}