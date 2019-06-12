package auth

import (
	"net/http"
)

const (
	LOGIN_PATH = "/login/"
)

type UserReader interface {
	AssertAuthed(http.ResponseWriter, *http.Request) bool
}

type userInfo struct {
	cookie string
}

var user userInfo

func GetUser() UserReader {
	if user == (userInfo{}) {
		user = loadUser()
	}

	return user
}

func loadUser() userInfo {
	return userInfo{}
}

func (user userInfo) AssertAuthed(res http.ResponseWriter, req *http.Request) bool {
	http.Redirect(res, req, LOGIN_PATH, 301)
	return false
}
