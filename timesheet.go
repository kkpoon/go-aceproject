package aceproject

import (
	"net/http"
	"strconv"
	"time"

	sling "gopkg.in/dghubble/sling.v1"
)

// TimesheetSaveWorkItemResponse represents saveworkitem response
type TimesheetSaveWorkItemResponse struct {
	Status  string         `json:"status"`
	Results []SaveWorkItem `json:"results"`
}

// SaveWorkItem represents logging time sheet entry to ACEProject
type SaveWorkItem struct {
	UserID          *int64  `url:"UserId,omitempty"`
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

// GetTimeReportParam represents gettimereport request parameter
type GetTimeReportParam struct {
	View                    int    `url:"View"`
	FilterMyWorkItems       bool   `url:"FilterMyWorkItems"`
	FilterTimeCreatorUserID string `url:"FilterTimeCreatorUserId,omitempty"`
	FilterDateFrom          string `url:"FilterDateFrom,omitempty"`
	FilterDateTo            string `url:"FilterDateTo,omitempty"`
	ProjectID               string `url:"ProjectId,omitempty"`
}

// DailyTimeReportResponse represents daily time sheet listing response
type DailyTimeReportResponse struct {
	Status  string           `json:"status"`
	Results []DailyTimesheet `json:"results"`
}

// DailyTimesheet is representing daily time sheet in ACEProject
type DailyTimesheet struct {
	DateWorked       string  `json:"DATE_WORKED"`
	TimesheetLineID  int64   `json:"TIMESHEET_LINE_ID"`
	CompanyID        int64   `json:"COMPANY_ID"`
	ProjectID        int64   `json:"PROJECT_ID"`
	ProjectType      string  `json:"PROJECT_TYPE_NAME"`
	ProjectNumber    string  `json:"PROJECT_NUMBER"`
	ProjectName      string  `json:"PROJECT_NAME"`
	TaskID           int64   `json:"TASK_ID"`
	TaskNumber       float64 `json:"TASK_NUMBER"`
	TaskName         string  `json:"TASK_RESUME"`
	TaskGroupName    string  `json:"TASK_GROUP_NAME"`
	TaskType         string  `json:"TASK_TYPE_NAME"`
	TimeTypeID       int64   `json:"TIME_TYPE_ID"`
	TimeTypeName     string  `json:"TIME_TYPE_NAME"`
	DateCreated      string  `json:"DATE_CREATED"`
	DateModified     string  `json:"DATE_MODIFIED"`
	DateSubmitted    string  `json:"DATE_SUBMITTED"`
	Comment          string  `json:"COMMENT"`
	CreatorUserID    int64   `json:"CREATOR_USER_ID"`
	CreatorUsername  string  `json:"CREATOR_USERNAME"`
	UserGroupName    string  `json:"USER_GROUP_NAME"`
	ClientNumber     string  `json:"CLIENT_NUMBER"`
	ClientName       string  `json:"CLIENT_NAME"`
	ApprovalLevel    int     `json:"APPROVAL_LEVEL"`
	TimeStatusName   string  `json:"TIME_STATUS_NAME"`
	ApprovalDate     string  `json:"DATE_APPROVAL_DATE"`
	ApprovalUserID   int64   `json:"APPROVAL_USER_ID"`
	ApprovalUsername string  `json:"APPROVAL_USERNAME"`
	Total            float64 `json:"TOTAL"`
	ErrorDesc        *string `json:"ERRORDESCRIPTION,omitempty"`
}

// TimesheetService provides methods to interact with project specific action
type TimesheetService struct {
	sling *sling.Sling
}

// NewTimesheetService return a new TimesheetService
func NewTimesheetService(httpClient *http.Client, guidInfo *GUIDInfo) *TimesheetService {
	return &TimesheetService{
		sling: sling.New().Client(httpClient).Base(baseURL).QueryStruct(guidInfo),
	}
}

// SaveWorkItem saves the time sheet item to ACE Project
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

// ListAllDailyWithProject returns the list of all accessible daily time sheets of a project
func (s *TimesheetService) ListAllDailyWithProject(projectID int64) ([]DailyTimesheet, *http.Response, error) {
	resObj := new(DailyTimeReportResponse)
	resp, err := s.sling.New().
		QueryStruct(CreateFunctionParam("gettimereport")).
		QueryStruct(&GetTimeReportParam{
			View:              1,
			FilterMyWorkItems: false,
			ProjectID:         strconv.FormatInt(projectID, 10),
		}).
		ReceiveSuccess(resObj)
	if resObj != nil && len(resObj.Results) > 0 {
		if resObj.Results[0].ErrorDesc != nil {
			return nil, resp, Error{*resObj.Results[0].ErrorDesc}
		}
		return *(&resObj.Results), resp, err
	}
	return make([]DailyTimesheet, 0), resp, err
}

// ListAllDailyWithDateRange returns the list of all accessible daily time sheets within the date range
func (s *TimesheetService) ListAllDailyWithDateRange(from, to time.Time) ([]DailyTimesheet, *http.Response, error) {
	resObj := new(DailyTimeReportResponse)
	resp, err := s.sling.New().
		QueryStruct(CreateFunctionParam("gettimereport")).
		QueryStruct(&GetTimeReportParam{
			View:              1,
			FilterMyWorkItems: false,
			FilterDateFrom:    from.Format("2006-01-02"),
			FilterDateTo:      to.Format("2006-01-02"),
		}).
		ReceiveSuccess(resObj)
	if resObj != nil && len(resObj.Results) > 0 {
		if resObj.Results[0].ErrorDesc != nil {
			return nil, resp, Error{*resObj.Results[0].ErrorDesc}
		}
		return *(&resObj.Results), resp, err
	}
	return make([]DailyTimesheet, 0), resp, err
}

// ListAllDailyWithUserIDDateRange returns the of all daily time sheets of a user within the date range
func (s *TimesheetService) ListAllDailyWithUserIDDateRange(userID int64, from, to time.Time) ([]DailyTimesheet, *http.Response, error) {
	resObj := new(DailyTimeReportResponse)
	resp, err := s.sling.New().
		QueryStruct(CreateFunctionParam("gettimereport")).
		QueryStruct(&GetTimeReportParam{
			View:                    1,
			FilterMyWorkItems:       false,
			FilterTimeCreatorUserID: strconv.FormatInt(userID, 10),
			FilterDateFrom:          from.Format("2006-01-02"),
			FilterDateTo:            to.Format("2006-01-02"),
		}).
		ReceiveSuccess(resObj)
	if resObj != nil && len(resObj.Results) > 0 {
		if resObj.Results[0].ErrorDesc != nil {
			return nil, resp, Error{*resObj.Results[0].ErrorDesc}
		}
		return *(&resObj.Results), resp, err
	}
	return make([]DailyTimesheet, 0), resp, err
}
