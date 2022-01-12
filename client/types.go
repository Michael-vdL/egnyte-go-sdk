package client

/**
  Authentication
*/

type AuthResponse struct {
  AccessToken string `json:"access_token"`
  TokenType string `json:"token_type"`
  ExpiresIn int `json:"expires_in"`
}


/**
  Events
  - Cursor
  - EventsAtCursor
*/

type Cursor struct {
  Timestamp string `json:"timestamp"`
  LatestEventId int `json:"latest_event_id"`
  OldestEventId int `json:"oldest_event_id"`
}

type Events struct {
  Count int `json:"count"`
  Events []Event `json:"events"`
  LatestId int `json:"latest_id"`
  OldestId int `json:"oldest_id"`
}

type Event struct {
  Id int `json:"id"`
  Timestamp string `json:"timestamp"`
  Actor int `json:"actor"`
  Type string `json:"type"`
  Action string `json:"action"`
  EventData EventData `json:"data"`
  ActionSource string `json:"action_source"`
  ObjectDetail string `json:"object_detail"`
}

type EventData struct {
  TargetPath string `json:"target_path"`
  TargetId string `json:"target_id"`
  TargetGroupId string `json:"target_group_id"`
  IsFolder bool `json:"is_folder"`
}

/**
  Users
  - Get All Users (Paginated)
    - Query Parameters:
      - startIndex (optional) Non-negative integer >= 1
      - count (optional) Non-negative integer <= 100
  - Get User By ID


*/

type Users struct {
	StartIndex   int    `json:"startIndex"`
	TotalResults int    `json:"totalResults"`
	ItemsPerPage int    `json:"itemsPerPage"`
	Resources    []User `json:"resources"`
}

type User struct {
	ID                   string  `json:"id"`
	UserName             string  `json:"userName"`
	ExternalId           string  `json:"externalId"`
	Name                 Name    `json:"name"`
	Active               bool    `json:"active"`
	Locked               bool    `json:"locked"`
	EmailChangePending   bool    `json:"emailChangePending"`
	AuthType             string  `json:"authType"`
	UserType             string  `json:"userType"`
	Role                 string  `json:"role"`
	IdpUserId            string  `json:"idpUserId"`
	UserPrincipleName    string  `json:"userPrincipleName"`
	ExpiryDate           string  `json:"expiryDate"`
	DeleteOnExpirty      bool    `json:"deleteOnExpirty"`
	CreatedDate          string  `json:"createdDate"`
	LastModificationDate string  `json:"lastModificationDate"`
	LastActiveDate       string  `json:"lastActiveDate"`
	Groups               []Group `json:"groups"`
	IsServiceAccount     bool    `json:"isServiceAccount"`
}

type Name struct {
	Formatted  string `json:"formatted"`
	FamilyName string `json:"familyName"`
	GivenName  string `json:"givenName"`
}

type Group struct {
	DisplayName string `json:"displayName"`
	Value       string `json:"value"`
}
