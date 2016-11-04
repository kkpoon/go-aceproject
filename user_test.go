package aceproject_test

import (
	"net/http"
	"os"
	"testing"

	aceproject "github.com/kkpoon/go-aceproject"
)

func TestUserList(t *testing.T) {
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

	if guidInfo == nil {
		t.Error("Expected to login success")
	}
	if err != nil {
		t.Error("Expected no error, got ", err)
	}

	userSvc := aceproject.NewUserService(client, guidInfo)

	users, _, err := userSvc.List()

	if users == nil {
		t.Error("Expected to have a user list, but it is nil")
	} else if len(users) == 0 {
		t.Error("Expected to have a user list, but size=", len(users))
	}
	if err != nil {
		t.Error("Expected no error, got ", err)
	}
}

func TestUserListWithActiveness(t *testing.T) {
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

	if guidInfo == nil {
		t.Error("Expected to login success")
	}
	if err != nil {
		t.Error("Expected no error, got ", err)
	}

	userSvc := aceproject.NewUserService(client, guidInfo)

	users, _, err := userSvc.ListWithActiveness(true)

	if users == nil {
		t.Error("Expected to have a user list, but it is nil")
	} else if len(users) == 0 {
		t.Error("Expected to have a user list, but size=", len(users))
	}
	if err != nil {
		t.Error("Expected no error, got ", err)
	}
}
