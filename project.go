package aceproject

import (
	"net/http"

	"github.com/dghubble/sling"
)

// ProjectResponse represents porject listing response
type ProjectResponse struct {
	Status  string    `json:"status"`
	Results []Project `json:"results"`
}

// Project is representing project in ACEProject
type Project struct {
	ID            int64   `json:"PROJECT_ID"`
	Name          string  `json:"PROJECT_NAME"`
	ProjectNumber string  `json:"PROJECT_NUMBER"`
	TypeID        int64   `json:"PROJECT_TYPE"`
	Type          string  `json:"PROJECT_TYPE_NAME"`
	ErrorDesc     *string `json:"ERRORDESCRIPTION,omitempty"`
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
	resp, err := s.sling.New().
		QueryStruct(CreateFunctionParam("getprojects")).
		ReceiveSuccess(projRes)
	if projRes != nil && len(projRes.Results) > 0 {
		if projRes.Results[0].ErrorDesc != nil {
			return nil, resp, Error{*projRes.Results[0].ErrorDesc}
		}
		return *(&projRes.Results), resp, err
	}
	return make([]Project, 0), resp, err
}
