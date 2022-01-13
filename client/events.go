package client

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type EventClient struct {
  Client
}

func (c *EventClient) GetCursor() (*Cursor, error) {
  cursorUrl := fmt.Sprintf("%s%s/events/cursor", c.BaseUrl, c.APIVersion)

  req, err := http.NewRequest("GET", cursorUrl, nil)
  if err != nil {
    return nil, err
  }

  authHeader := fmt.Sprintf("Bearer %s", c.authToken)
  req.Header.Add("Authorization", authHeader)

  res, err := c.httpClient.Do(req)
  if err != nil {
    return nil, err
  }
  defer res.Body.Close()
  if res.StatusCode != 200 {
    return nil, fmt.Errorf("get cursor failed. status code: %d", res.StatusCode)
  }

  var cursor Cursor
  err = json.NewDecoder(res.Body).Decode(&cursor)
  if err != nil {
    return nil, err
  }

  return &cursor, nil
}

func (c *EventClient) GetEvents(cursor string) (*Events, error) {
  // TODO - Once Query Params are Implement, add cursor and type to Query Param Handler
  typeQueryParam := "type=file_system|note|permission_change"
  eventsUrl := fmt.Sprintf("%s%s/events?id=%s&%s", c.BaseUrl, c.APIVersion, cursor, typeQueryParam)

  req, err := http.NewRequest("GET", eventsUrl, nil)
  if err != nil {
    return nil, err
  }

  authHeader := fmt.Sprintf("Bearer %s", c.authToken)
  req.Header.Add("Authorization", authHeader)

  res, err := c.httpClient.Do(req)
  if err != nil {
    return nil, err
  }
  defer res.Body.Close()
  if res.StatusCode != 200 {
    return nil, fmt.Errorf("get events failed. status code: %d", res.StatusCode)
  }

  var events Events
  err = json.NewDecoder(res.Body).Decode(&events)
  if err != nil {
    return nil, err
  }

  return &events, nil
}