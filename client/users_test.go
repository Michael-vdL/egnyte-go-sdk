package client

import (
	"testing"
)

func Test_GetUsers(t *testing.T) {
  // Setup
  config, _ := ConfigurationFromFile("../config.json")
  c := NewClient(config.SubDomain, config.APIVersion)
  c.Authenticate(config.Username, config.Password, config.APIKey, config.APISecret)
  uc := &UserClient{Client: *c}

  // Test
  users, err := uc.GetUsers()
  if err != nil {
    t.Errorf("GetUsers failed: %s", err)
  }
  
  t.Log(users)
 
}