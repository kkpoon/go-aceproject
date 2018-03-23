package aceproject

import (
	"net/http"
	"strconv"

	sling "gopkg.in/dghubble/sling.v1"
)

// GetProjectsParam represents getprojects request parameter
type GetProjectsParam struct {
	FilterCompletedProject *bool  `url:"Filtercompletedproject,omitempty"`
	ProjectID              string `url:"Projectid,omitempty"`
}

// ProjectResponse represents porject listing response
type ProjectResponse struct {
	Status  string    `json:"status"`
	Results []Project `json:"results"`
}

// Project is representing project in ACEProject
type Project struct {
	ID                 int64   `json:"PROJECT_ID"`
	Name               string  `json:"PROJECT_NAME"`
	ProjectNumber      string  `json:"PROJECT_NUMBER"`
	TypeID             int64   `json:"PROJECT_TYPE"`
	Type               string  `json:"PROJECT_TYPE_NAME"`
	ProjectStatusName  string  `json:"PROJECT_STATUS_NAME"`
	PriorityName       string  `json:"PROJECT_PRIORITY_NAME"`
	ClientName         string  `json:"CLIENT_NAME"`
	ContactName        string  `json:"CONTACT_NAME"`
	ContactPhone       string  `json:"CONTACT_PHONE"`
	DateCreated        string  `json:"DATE_CREATED"`
	DateModified       string  `json:"DATE_MODIFIED"`
	ProjectCreatorID   int64   `json:"PROJECT_CREATOR_ID"`
	Username           string  `json:"USERNAME"`
	UserUpdateID       int64   `json:"USER_UPDATE_ID"`
	UserUpdateUsername string  `json:"USER_UPDATE_USERNAME"`
	AssignedUsers      string  `json:"ASSIGNED_USERS"`
	ErrorDesc          *string `json:"ERRORDESCRIPTION,omitempty"`
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

// ListWithCompleteness returns the list of complete / incomplete projects
func (s *ProjectService) ListWithCompleteness(complete bool) ([]Project, *http.Response, error) {
	projRes := new(ProjectResponse)
	resp, err := s.sling.New().
		QueryStruct(CreateFunctionParam("getprojects")).
		QueryStruct(&GetProjectsParam{FilterCompletedProject: &complete}).
		ReceiveSuccess(projRes)
	if projRes != nil && len(projRes.Results) > 0 {
		if projRes.Results[0].ErrorDesc != nil {
			return nil, resp, Error{*projRes.Results[0].ErrorDesc}
		}
		return *(&projRes.Results), resp, err
	}
	return make([]Project, 0), resp, err
}

// Get returns the project of given ID
func (s *ProjectService) Get(projectID int64) (*Project, *http.Response, error) {
	//
	projRes := new(ProjectResponse)
	resp, err := s.sling.New().
		QueryStruct(CreateFunctionParam("getprojects")).
		QueryStruct(&GetProjectsParam{ProjectID: strconv.FormatInt(projectID, 10)}).
		ReceiveSuccess(projRes)
	if projRes != nil && len(projRes.Results) > 0 {
		if projRes.Results[0].ErrorDesc != nil {
			return nil, resp, Error{*projRes.Results[0].ErrorDesc}
		}
		return &projRes.Results[0], resp, err
	}
	return nil, resp, err
}
