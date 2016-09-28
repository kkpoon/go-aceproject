package aceproject

import (
	"net/http"

	"github.com/dghubble/sling"
)

type ProjectResponse struct {
	Status  string    `json:"status"`
	Results []Project `json:"results"`
}

// Project is representing project in ACEProject
type Project struct {
	ID int `json:"ID"`
}

// ProjectService provides methods to interact with project specific action
type ProjectService struct {
	sling *sling.Sling
}

// NewProjectService return a new ProjectService
func NewProjectService(httpClient *http.Client, guidInfo *GUIDInfo) *ProjectService {
	return &ProjectService{
		sling: sling.New().Client(httpClient).Base(baseURL).QueryStruct(guidInfo),
	}
}

// List returns the project list
func (s *ProjectService) List() ([]Project, *http.Response, error) {
	projRes := new(ProjectResponse)
	resp, err := s.sling.New().QueryStruct(CreateFunctionParam("getprojects")).ReceiveSuccess(projRes)
	if projRes != nil && len(projRes.Results) > 0 {
		return *(&projRes.Results), resp, err
	}
	return make([]Project, 0), resp, err
}
