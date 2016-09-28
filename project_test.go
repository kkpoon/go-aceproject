package aceproject_test

import (
	"net/http"
	"os"
	"testing"

	aceproject "github.com/kkpoon/go-aceproject"
)

func TestProjectList(t *testing.T) {
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

	if guidInfo == nil && len((*guidInfo).GUID) != 36 {
		t.Error("Expected to login success")
	}
	if err != nil {
		t.Error("Expected no error, got ", err)
	}

	proj := aceproject.NewProjectService(client, guidInfo)

	projects, _, err := proj.List()

	if projects == nil && len(projects) > 0 {
		t.Error("Expected to have a project list, but size=", len(projects))
	}
	if err != nil {
		t.Error("Expected no error, got ", err)
	}
}
