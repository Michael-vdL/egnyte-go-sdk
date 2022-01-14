package client

import (
	"testing"
)

func Test_CreateFolder(t *testing.T) {
  // Setup
  config, _ := ConfigurationFromFile("../config.json")
  c := NewClient(config.SubDomain, config.APIVersion)
  c.Authenticate(config.Username, config.Password, config.APIKey, config.APISecret)
  fc := &FSClient{Client: *c}

  // Test
  createFolderResponse, err := fc.CreateFolder("Shared/test")
  if err != nil {
    t.Errorf("CreateFolder failed: %s", err)
  }
  t.Log(createFolderResponse)
}