package aceproject_test

import (
	"net/http"
	"os"
	"testing"
	"time"

	aceproject "github.com/kkpoon/go-aceproject"
)

func TestTaskList(t *testing.T) {
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
		taskSvc := aceproject.NewTaskService(client, guidInfo)

		tasks, _, err := taskSvc.List()

		if tasks == nil {
			t.Error("Expected to have a task list, but it is nil")
		} else if len(tasks) == 0 {
			t.Error("Expected to have a task list, but size=", len(tasks))
		}
		if err != nil {
			t.Error("Expected no error, got ", err)
		}
	}
}

func TestTaskListWithProjectId(t *testing.T) {
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

		taskSvc := aceproject.NewTaskService(client, guidInfo)

		projectID := projects[0].ID
		tasks, _, err := taskSvc.ListWithProject(projectID)

		if tasks == nil {
			t.Error("Expected to have a task list, but it is nil")
		} else if len(tasks) == 0 {
			t.Error("Expected to have a task list, but size=", len(tasks))
		}
		if err != nil {
			t.Error("Expected no error, got ", err)
		}
	}
}

func TestTaskListWithLastModifiedDate(t *testing.T) {
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
		taskSvc := aceproject.NewTaskService(client, guidInfo)

		tasks, _, err := taskSvc.ListWithLastModifiedDate(time.Now())

		if tasks == nil {
			t.Error("Expected to have a task list, but it is nil")
		} else if len(tasks) == 0 {
			t.Error("Expected to have a task list, but size=", len(tasks))
		}
		if err != nil {
			t.Error("Expected no error, got ", err)
		}
	}
}
