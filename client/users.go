package client

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type UserClient struct {
	Client 
}

func (c *UserClient) GetUsers() (*Users, error) {
	usersUrl := fmt.Sprintf("%s/%s/users", c.BaseUrl, c.APIVersion)

	req, err := http.NewRequest("GET", usersUrl, nil)
  if err != nil {
    return nil, err
  }
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.authToken))

  res, err := c.httpClient.Do(req)
  if err != nil {
    return nil, err
  }
  defer res.Body.Close()
  if res.StatusCode != 200 {
    return nil, fmt.Errorf("get users failed. status code: %d", res.StatusCode)
  }

  var users Users
  err = json.NewDecoder(res.Body).Decode(&users)
  if err != nil {
    return nil, err
  }

	return &users, nil
}

func (c *UserClient) GetUser(userId string) (*User, error) {
	return nil, nil
}