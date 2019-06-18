package auth

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/multiplio/ozymandias/version"
)

const (
	login_path  = "/login"
	config_path = "~/.config/" + version.AppName + "/user"
)

type UserReader interface {
	AssertAuthed(http.ResponseWriter, *http.Request) bool
}

type userInfo struct {
	Cookie string
	Token  string
}

var user userInfo

func GetUser() UserReader {
	if user == (userInfo{}) {
		user = loadUser()
	}

	return user
}

func loadUser() userInfo {
	content, err := ioutil.ReadFile(config_path)
	if err != nil {
		return userInfo{}
	}

	var config userInfo
	err = json.Unmarshal(content, &config)
	if err != nil {
		return userInfo{}
	}

	return config
}

func (u userInfo) saveUser() error {
	data, err := json.Marshal(u)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(config_path, data, 0600)
}

func (u userInfo) AssertAuthed(res http.ResponseWriter, req *http.Request) bool {
	http.Redirect(res, req, login_path, http.StatusTemporaryRedirect)
	return false
}
