package aceproject_test

import (
	"net/http"
	"os"
	"testing"

	aceproject "github.com/kkpoon/go-aceproject"
)

func TestLoginSuccess(t *testing.T) {
	accountid := os.Getenv("ACE_ACCOUNTID")
	username := os.Getenv("ACE_USERNAME")
	password := os.Getenv("ACE_PASSWORD")

	if accountid == "" || username == "" || password == "" {
		t.Error("ACE_ACCOUNTID, ACE_USERNAME, ACE_PASSWORD are not set in environment variable")
	}

	authInfo := aceproject.AuthInfo{accountid, username, password}
	svc := aceproject.NewLoginService(&http.Client{})
	guidInfo, _, err := svc.Login(&authInfo)

	if guidInfo == nil {
		t.Error("Expected to get GUID, got no GUID")
	} else if len((*guidInfo).GUID) != 36 {
		t.Error("Expected to get GUID, got GUID with length of ", len((*guidInfo).GUID))
	}
	if err != nil {
		t.Error("Expected no error, got ", err)
	}
}

func TestLoginFail(t *testing.T) {
	authInfo := aceproject.AuthInfo{"", "", ""}
	svc := aceproject.NewLoginService(&http.Client{})
	guidInfo, _, err := svc.Login(&authInfo)

	if guidInfo != nil {
		t.Error("Expected to get nil GUID, got ", guidInfo)
	}
	if err == nil {
		t.Error("Expected it has err, got ", err)
	}
}
