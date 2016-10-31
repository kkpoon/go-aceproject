package aceproject

import (
	"net/http"

	"github.com/dghubble/sling"
)

// TimesheetGetMyWeeksResponse represents GetMyWeeks response
type TimesheetGetMyWeeksResponse struct {
	Status  string    `json:"status"`
	Results []MyWeeks `json:"results"`
}

// MyWeeks represents MyWeeks in ACEProject
type MyWeeks struct {
	ID int `json:"ID"`
}

// TimesheetService provides methods to interact with project specific action
type TimesheetService struct {
	sling *sling.Sling
}

// NewTimesheetService return a new ProjectService
func NewTimesheetService(httpClient *http.Client, guidInfo *GUIDInfo) *TimesheetService {
	return &TimesheetService{
		sling: sling.New().Client(httpClient).Base(baseURL).QueryStruct(guidInfo),
	}
}

// GetMyWeeks returns the project list
func (s *TimesheetService) GetMyWeeks() ([]MyWeeks, *http.Response, error) {
	res := new(TimesheetGetMyWeeksResponse)
	httpResp, err := s.sling.New().
		QueryStruct(CreateFunctionParam("getmyweeks")).
		ReceiveSuccess(res)
	if res != nil && len(res.Results) > 0 {
		return *(&res.Results), httpResp, err
	}
	return make([]MyWeeks, 0), httpResp, err
}
