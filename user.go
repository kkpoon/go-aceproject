package aceproject

import (
	"net/http"

	sling "gopkg.in/dghubble/sling.v1"
)

// GetUsersParam represents getusers request parameter
type GetUsersParam struct {
	FilterActive bool `url:"FilterActive,omitempty"`
}

// UserResponse represents user listing response
type UserResponse struct {
	Status  string `json:"status"`
	Results []User `json:"results"`
}

// User is representing user in ACEProject
type User struct {
	ID            int64   `json:"USER_ID"`
	Username      string  `json:"USERNAME"`
	Email         string  `json:"EMAIL_ALERT"`
	FirstName     string  `json:"FIRST_NAME"`
	LastName      string  `json:"LAST_NAME"`
	UserGroupID   int64   `json:"USER_GROUP_ID"`
	UserGroupName string  `json:"USER_GROUP_NAME"`
	Active        bool    `json:"ACTIVE"`
	ErrorDesc     *string `json:"ERRORDESCRIPTION,omitempty"`
}

// UserService provides methods to interact with user specific action
type UserService struct {
	sling *sling.Sling
}

// NewUserService return a new UserService
func NewUserService(httpClient *http.Client, guidInfo *GUIDInfo) *UserService {
	return &UserService{
		sling: sling.New().Client(httpClient).Base(baseURL).QueryStruct(guidInfo),
	}
}

// List returns the user list
func (s *UserService) List() ([]User, *http.Response, error) {
	resObj := new(UserResponse)
	resp, err := s.sling.New().
		QueryStruct(CreateFunctionParam("getusers")).
		ReceiveSuccess(resObj)
	if resObj != nil && len(resObj.Results) > 0 {
		if resObj.Results[0].ErrorDesc != nil {
			return nil, resp, Error{*resObj.Results[0].ErrorDesc}
		}
		return *(&resObj.Results), resp, err
	}
	return make([]User, 0), resp, err
}

// ListWithActiveness returns the list of active / non-active users
func (s *UserService) ListWithActiveness(active bool) ([]User, *http.Response, error) {
	resObj := new(UserResponse)
	resp, err := s.sling.New().
		QueryStruct(CreateFunctionParam("getusers")).
		QueryStruct(&GetUsersParam{FilterActive: active}).
		ReceiveSuccess(resObj)
	if resObj != nil && len(resObj.Results) > 0 {
		if resObj.Results[0].ErrorDesc != nil {
			return nil, resp, Error{*resObj.Results[0].ErrorDesc}
		}
		return *(&resObj.Results), resp, err
	}
	return make([]User, 0), resp, err
}
