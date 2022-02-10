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
  queryParameters := make(map[string]string)
  queryParameters["count"] = "100"
  users, err := uc.GetUsers(queryParameters)
  if err != nil {
    t.Errorf("GetUsers failed: %s", err)
  }
  
  t.Log(users)
  t.Log(len(users.Resources))
}

func Test_GetUserById(t *testing.T) {
  // Setup
  config, _ := ConfigurationFromFile("../config.json")
  c := NewClient(config.SubDomain, config.APIVersion)
  c.Authenticate(config.Username, config.Password, config.APIKey, config.APISecret)
  uc := &UserClient{Client: *c}

  // Test
  user, err := uc.GetUserById("1") // Any/All Tenants should have a user with id 1
  if err != nil {
    t.Errorf("GetUserById failed: %s", err)
  }

  t.Log(user)
}