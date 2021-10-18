package aceproject

import (
	"fmt"
	"net/http"
)

const baseURL = "https://api.aceproject.com"

// Function represents the ACEProject fct parameter in URL
type Function struct {
	Fct    string `url:"fct"`
	Format string `url:"format"`
}

// Error represents the error return from ACEProject API
type Error struct {
	Message string `json:"ERRORDESCRIPTION"`
}

// ErrorResponse represents the error return from ACEProject API
type ErrorResponse struct {
	Status  string  `json:"status"`
	Results []Error `json:"results"`
}

func (e Error) Error() string {
	return fmt.Sprintf("aceproject: %v", e.Message)
}

// CreateFunctionParam creates the url query param for ACEProject function
func CreateFunctionParam(fct string) *Function {
	return &Function{fct, "JSON"}
}

// Client is a ACEProject RESTful API client
type Client struct {
	ProjectService *ProjectService
}

// NewClient creates new client for ACEProject API service
func NewClient(httpClient *http.Client, authInfo *AuthInfo) (*Client, error) {
	loginService := NewLoginService(httpClient)

	guidInfo, _, err := loginService.Login(authInfo)

	if err != nil {
		return nil, err
	}

	if guidInfo == nil {
		return nil, fmt.Errorf("aceproject API server does not return login info")
	}

	return &Client{
		ProjectService: NewProjectService(httpClient, guidInfo),
	}, nil
}
