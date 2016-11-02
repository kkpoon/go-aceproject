package aceproject

import (
	"net/http"

	"github.com/dghubble/sling"
)

// GetTasksParam represents gettasks request parameter
type GetTasksParam struct {
	ProjectID int64 `url:"ProjectID,omitempty"`
}

// TaskResponse represents task listing response
type TaskResponse struct {
	Status  string `json:"status"`
	Results []Task `json:"results"`
}

// Task is representing task in ACEProject
type Task struct {
	ID            int     `json:"TASK_ID"`
	Name          string  `json:"TASK_RESUME"`
	TaskGroupID   int     `json:"TASK_GROUP_ID"`
	TaskGroupName string  `json:"TASK_GROUP_NAME"`
	ProjectID     int     `json:"PROJECT_ID"`
	ProjectName   string  `json:"PROJECT_NAME"`
	ErrorDesc     *string `json:"ERRORDESCRIPTION,omitempty"`
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
		QueryStruct(&GetTasksParam{ProjectID: projectID}).
		ReceiveSuccess(resObj)
	if resObj != nil && len(resObj.Results) > 0 {
		if resObj.Results[0].ErrorDesc != nil {
			return nil, resp, Error{*resObj.Results[0].ErrorDesc}
		}
		return *(&resObj.Results), resp, err
	}
	return make([]Task, 0), resp, err
}
