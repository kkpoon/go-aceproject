package aceproject

import (
	"net/http"

	"github.com/dghubble/sling"
)

// AuthInfo is representing the information for authentication
type AuthInfo struct {
	AccountID string `url:"AccountId"`
	Username  string `url:"UserName"`
	Password  string `url:"Password"`
}

// GUIDInfo is representing the information for endpoint auth
type GUIDInfo struct {
	GUID string `url:"guid" json:"GUID,omitempty"`
}

// LoginResponse represents the success login response from ACEProject API
type LoginResponse struct {
	Status  string     `json:"status"`
	Results []GUIDInfo `json:"results"`
}

// LoginService provides methods to interact with login action
type LoginService struct {
	sling *sling.Sling
}

// Login performs login action to ACEProject API
func (s *LoginService) Login(params *AuthInfo) (*GUIDInfo, *http.Response, error) {
	loginRes := new(LoginResponse)
	resp, err := s.sling.New().QueryStruct(CreateFunctionParam("login")).QueryStruct(params).ReceiveSuccess(loginRes)
	if loginRes != nil && len(loginRes.Results) > 0 {
		return &loginRes.Results[0], resp, err
	}
	return nil, resp, err
}

// NewLoginService creates a new LoginService
func NewLoginService(httpClient *http.Client) *LoginService {
	return &LoginService{
		sling: sling.New().Client(httpClient).Base(baseURL),
	}
}
