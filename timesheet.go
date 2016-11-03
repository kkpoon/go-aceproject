package aceproject

import (
	"net/http"

	"github.com/dghubble/sling"
)

// TimesheetSaveWorkItemResponse represents GetMyWeeks response
type TimesheetSaveWorkItemResponse struct {
	Status  string         `json:"status"`
	Results []SaveWorkItem `json:"results"`
}

// SaveWorkItem represents logging timesheet entry to ACEProject
type SaveWorkItem struct {
	TimesheetLineID *int64  `url:"TimesheetLineId,omitempty" json:"TIMESHEET_LINE_ID,omitempty"`
	WeekStart       string  `url:"WeekStart,omitempty"`
	TaskID          int64   `url:"TaskId"`
	TimeTypeID      int64   `url:"TimetypeId"`
	HoursDay1       float64 `url:"HoursDay1"`
	HoursDay2       float64 `url:"HoursDay2"`
	HoursDay3       float64 `url:"HoursDay3"`
	HoursDay4       float64 `url:"HoursDay4"`
	HoursDay5       float64 `url:"HoursDay5"`
	HoursDay6       float64 `url:"HoursDay6"`
	HoursDay7       float64 `url:"HoursDay7"`
	Comments        *string `url:"Comments,omitempty"`
	ErrorDesc       *string `json:"ERRORDESCRIPTION,omitempty"`
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
func (s *TimesheetService) SaveWorkItem(item *SaveWorkItem) (*SaveWorkItem, *http.Response, error) {
	resObj := new(TimesheetSaveWorkItemResponse)
	httpResp, err := s.sling.New().
		QueryStruct(CreateFunctionParam("saveworkitem")).
		QueryStruct(item).
		ReceiveSuccess(resObj)
	if resObj != nil && len(resObj.Results) > 0 {
		if resObj.Results[0].ErrorDesc != nil {
			return nil, httpResp, Error{*resObj.Results[0].ErrorDesc}
		}
		return &resObj.Results[0], httpResp, err
	}
	return nil, httpResp, err
}
