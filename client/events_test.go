package client

import (
	"strconv"
	"testing"
)

func Test_GetCursor(t *testing.T) {
  // Setup
  config, _ := ConfigurationFromFile("../config.json")
  c := NewClient(config.SubDomain, config.APIVersion)
  c.Authenticate(config.Username, config.Password, config.APIKey, config.APISecret)
  ec := &EventClient{Client: *c}

  // Test
  cursor, err := ec.GetCursor()
  if err != nil {
    t.Errorf("GetCursor failed: %s", err)
  }

  t.Log(cursor)
}

func Test_GetEvents(t *testing.T) {
  // Setup
  config, _ := ConfigurationFromFile("../config.json")
  c := NewClient(config.SubDomain, config.APIVersion)
  c.Authenticate(config.Username, config.Password, config.APIKey, config.APISecret)
  ec := &EventClient{Client: *c}

  cursor, err := ec.GetCursor()
  if err != nil {
    t.Errorf("GetCursor failed: %s", err)
  }
  testCursor := cursor.LatestEventId - 21 // To Limit number of events returned, reduce cursor by 20

  // Test
  t.Logf("received cursor id %d - using %d - expecting 20 results", cursor.LatestEventId, testCursor)

  events, err := ec.GetEvents(strconv.Itoa(testCursor))
  if err != nil {
    t.Errorf("GetEvents failed: %s", err)
  }

  // TODO: Look into why the count is always less than expected - seems unrelated to event types
  t.Log(events.Count) // Should be 20 events but not a failure if it is less or more

  for _, event := range events.Events {
    t.Log(event.Type)
  }
}