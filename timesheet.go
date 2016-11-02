package aceproject

import (
	"net/http"

	"github.com/dghubble/sling"
)

// TimesheetSaveWorkItemResponse represents GetMyWeeks response
type TimesheetSaveWorkItemResponse struct {
	Status  string               `json:"status"`
	Results []SaveWorkItemResult `json:"results"`
}

// SaveWorkItem represents logging timesheet entry to ACEProject
type SaveWorkItem struct {
	WeekStart  string  `url:"WeekStart,omitempty"`
	TaskID     int64   `url:"TaskId"`
	TimeTypeID int64   `url:"TimetypeId"`
	HoursDay1  float64 `url:"HoursDay1,omitempty"`
	HoursDay2  float64 `url:"HoursDay2,omitempty"`
	HoursDay3  float64 `url:"HoursDay3,omitempty"`
	HoursDay4  float64 `url:"HoursDay4,omitempty"`
	HoursDay5  float64 `url:"HoursDay5,omitempty"`
	HoursDay6  float64 `url:"HoursDay6,omitempty"`
	HoursDay7  float64 `url:"HoursDay7,omitempty"`
	Comments   *string `url:"Comments,omitempty"`
}

// SaveWorkItemResult represents response of timesheet entry to ACEProject
type SaveWorkItemResult struct {
	ErrorDesc *string `json:"ERRORDESCRIPTION,omitempty"`
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

// SaveWorkItem saves the timesheet item to ACE Project
func (s *TimesheetService) SaveWorkItem(item *SaveWorkItem) (*http.Response, error) {
	resObj := new(TimesheetSaveWorkItemResponse)
	httpResp, err := s.sling.New().
		QueryStruct(CreateFunctionParam("saveworkitem")).
		QueryStruct(item).
		ReceiveSuccess(resObj)
	if resObj != nil && len(resObj.Results) > 0 {
		if resObj.Results[0].ErrorDesc != nil {
			return httpResp, Error{*resObj.Results[0].ErrorDesc}
		}
		return httpResp, err
	}
	return httpResp, err
}
