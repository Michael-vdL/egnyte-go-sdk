package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type FSClient struct {
  Client
}

type CreateFolder struct {
  Action string `json:"action"`
}

func (c *FSClient) CreateFolder(path string) (*CreateFolderResponse, error) {
  // Only supports v1 API
  createFolderUrl := fmt.Sprintf("%s/pubapi/v1/fs/%s", c.BaseUrl, path)
  log.Default().Println(createFolderUrl)

  actionData := new(bytes.Buffer)
  action := CreateFolder{Action: "add_folder"}
  json.NewEncoder(actionData).Encode(&action)

  req, err := http.NewRequest("POST", createFolderUrl, actionData)
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
    return nil, fmt.Errorf("create folder failed. status code: %d", res.StatusCode)
  }

  var createFolderResponse CreateFolderResponse
  err = json.NewDecoder(res.Body).Decode(&createFolderResponse)
  if err != nil {
    return nil, err
  }

  return &createFolderResponse, nil
}