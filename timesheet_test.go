package aceproject_test

import (
	"net/http"
	"os"
	"testing"
	"time"

	aceproject "github.com/kkpoon/go-aceproject"
)

func TestDailyTimesheetListWithProjectId(t *testing.T) {
	accountid := os.Getenv("ACE_ACCOUNTID")
	username := os.Getenv("ACE_USERNAME")
	password := os.Getenv("ACE_PASSWORD")

	if accountid == "" || username == "" || password == "" {
		t.Error("ACE_ACCOUNTID, ACE_USERNAME, ACE_PASSWORD are not set in environment variable")
	}

	authInfo := aceproject.AuthInfo{accountid, username, password}
	client := &http.Client{}
	svc := aceproject.NewLoginService(client)
	guidInfo, _, err := svc.Login(&authInfo)

	if err != nil {
		t.Error("Expected no error, got ", err)
	}
	if guidInfo == nil {
		t.Error("Expected to login success")
	} else {
		proj := aceproject.NewProjectService(client, guidInfo)

		projects, _, err := proj.List()

		if projects == nil {
			t.Error("Expected to have a project list, but it is nil")
		} else if len(projects) == 0 {
			t.Error("Expected to have a project list, but size=", len(projects))
		}
		if err != nil {
			t.Error("Expected no error, got ", err)
		}

		tsSvc := aceproject.NewTimesheetService(client, guidInfo)

		projectID := projects[0].ID
		ts, _, err := tsSvc.ListAllDailyWithProject(projectID)

		if ts == nil {
			t.Error("Expected to have a daily time sheet list, but it is nil")
		}
		if err != nil {
			t.Error("Expected no error, got ", err)
		}
	}
}

func TestDailyTimesheetListWithDateRange(t *testing.T) {
	accountid := os.Getenv("ACE_ACCOUNTID")
	username := os.Getenv("ACE_USERNAME")
	password := os.Getenv("ACE_PASSWORD")

	if accountid == "" || username == "" || password == "" {
		t.Error("ACE_ACCOUNTID, ACE_USERNAME, ACE_PASSWORD are not set in environment variable")
	}

	authInfo := aceproject.AuthInfo{accountid, username, password}
	client := &http.Client{}
	svc := aceproject.NewLoginService(client)
	guidInfo, _, err := svc.Login(&authInfo)

	if err != nil {
		t.Error("Expected no error, got ", err)
	}
	if guidInfo == nil {
		t.Error("Expected to login success")
	} else {
		tsSvc := aceproject.NewTimesheetService(client, guidInfo)

		ts, _, err := tsSvc.ListAllDailyWithDateRange(
			time.Date(2016, 10, 1, 0, 0, 0, 0, time.UTC),
			time.Date(2016, 10, 6, 0, 0, 0, 0, time.UTC),
		)

		if ts == nil {
			t.Error("Expected to have a daily time sheet list, but it is nil")
		}
		if err != nil {
			t.Error("Expected no error, got ", err)
		}
	}
}
