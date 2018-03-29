package aceproject

import (
	"net/http"
	"time"

	sling "gopkg.in/dghubble/sling.v1"
)

// GetTasksParam represents gettasks request parameter
type GetTasksParam struct {
	ProjectID               *int64  `url:"ProjectID,omitempty"`
	FilterFirstDate         *string `url:"FilterFirstDate,omitempty"`
	FilterFirstDateOperator *int    `url:"FilterFirstDateOperator,omitempty"`
	FilterFirstDateValue    *string `url:"FilterFirstDateValue,omitempty"`
}

// TaskResponse represents task listing response
type TaskResponse struct {
	Status  string `json:"status"`
	Results []Task `json:"results"`
}

// Task is representing task in ACEProject
type Task struct {
	ID               int64   `json:"TASK_ID"`
	Name             string  `json:"TASK_RESUME"`
	TaskGroupID      int64   `json:"TASK_GROUP_ID"`
	TaskGroupName    string  `json:"TASK_GROUP_NAME"`
	TaskStatusName   string  `json:"TASK_STATUS_NAME"`
	ProjectID        int64   `json:"PROJECT_ID"`
	ProjectName      string  `json:"PROJECT_NAME"`
	DateTaskCreated  string  `json:"DATE_TASK_CREATED"`
	DateTaskModified string  `json:"DATE_TASK_MODIFIED"`
	UserCreatorID    int64   `json:"USER_CREATOR_ID"`
	Username         string  `json:"USERNAME"`
	UpdateUserID     int64   `json:"UPDATE_USER_ID"`
	UpdateUsername   string  `json:"UPDATE_USERNAME"`
	Assigned         string  `json:"ASSIGNED"`
	ErrorDesc        *string `json:"ERRORDESCRIPTION,omitempty"`
}

// TaskService provides methods to interact with task specific action
type TaskService struct {
	sling *sling.Sling
}

// NewTaskService return a new TaskService
func NewTaskService(httpClient *http.Client, guidInfo *GUIDInfo) *TaskService {
	return &TaskService{
		sling: sling.New().Client(httpClient).Base(baseURL).QueryStruct(guidInfo),
	}
}

// List returns all tasks
func (s *TaskService) List() ([]Task, *http.Response, error) {
	resObj := new(TaskResponse)
	resp, err := s.sling.New().
		QueryStruct(CreateFunctionParam("gettasks")).
		ReceiveSuccess(resObj)
	if resObj != nil && len(resObj.Results) > 0 {
		if resObj.Results[0].ErrorDesc != nil {
			return nil, resp, Error{*resObj.Results[0].ErrorDesc}
		}
		return *(&resObj.Results), resp, err
	}
	return make([]Task, 0), resp, err
}

// ListWithProject returns all tasks of a project
func (s *TaskService) ListWithProject(projectID int64) ([]Task, *http.Response, error) {
	resObj := new(TaskResponse)
	resp, err := s.sling.New().
		QueryStruct(CreateFunctionParam("gettasks")).
		QueryStruct(&GetTasksParam{ProjectID: &projectID}).
		ReceiveSuccess(resObj)
	if resObj != nil && len(resObj.Results) > 0 {
		if resObj.Results[0].ErrorDesc != nil {
			return nil, resp, Error{*resObj.Results[0].ErrorDesc}
		}
		return *(&resObj.Results), resp, err
	}
	return make([]Task, 0), resp, err
}

// ListWithLastModifiedDate returns all tasks modified on or after the specified date
func (s *TaskService) ListWithLastModifiedDate(modifiedDate time.Time) ([]Task, *http.Response, error) {
	filterDateType := "DATE_TASK_MODIFIED"
	filterDateOperator := 3 // Is Later Than Or On
	modifiedDateStr := modifiedDate.Format("2006-01-02")
	resObj := new(TaskResponse)
	resp, err := s.sling.New().
		QueryStruct(CreateFunctionParam("gettasks")).
		QueryStruct(&GetTasksParam{
			FilterFirstDate:         &filterDateType,
			FilterFirstDateOperator: &filterDateOperator,
			FilterFirstDateValue:    &modifiedDateStr,
		}).
		ReceiveSuccess(resObj)
	if resObj != nil && len(resObj.Results) > 0 {
		if resObj.Results[0].ErrorDesc != nil {
			return nil, resp, Error{*resObj.Results[0].ErrorDesc}
		}
		return *(&resObj.Results), resp, err
	}
	return make([]Task, 0), resp, err
}
