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
	usersUrl := fmt.Sprintf("%s%s/users", c.BaseUrl, c.APIVersion)

	req, err := http.NewRequest("GET", usersUrl, nil)
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
    return nil, fmt.Errorf("get users failed. status code: %d", res.StatusCode)
  }

  var users Users
  err = json.NewDecoder(res.Body).Decode(&users)
  if err != nil {
    return nil, err
  }

	return &users, nil
}

func (c *UserClient) GetUserById(userId string) (*User, error) {
  userByIdUrl := fmt.Sprintf("%s%s/users/%s", c.BaseUrl, c.APIVersion, userId)
  
  req, err := http.NewRequest("GET", userByIdUrl, nil)
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
    return nil, fmt.Errorf("get user by id failed. status code: %d", res.StatusCode)
  }

  var user User
  err = json.NewDecoder(res.Body).Decode(&user)
  if err != nil {
    return nil, err
  }

  return &user, nil
}