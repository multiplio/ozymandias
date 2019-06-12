package auth

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/google/go-github/v25/github"

	"github.com/multiplio/ozymandias/version"
)

var OAuthAppURL = "https://multiplio.github.io/ozymandias/"

// Handler returns http.Handler for API endpoint
func Handler() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "application/json")

		body := "login"

		res.WriteHeader(200)
		res.Write([]byte(body))
	}
}

func FindOrCreateToken(user, password, twoFactorCode string) (token string, err error) {
	if len(password) >= 40 && isToken(password) {
		return password, nil
	}

	transport := github.BasicAuthTransport{
		Username: user,
		Password: password,
	}
	if twoFactorCode != "" {
		transport.OTP = twoFactorCode
	}
	client := github.NewClient(transport.Client())

	request := github.AuthorizationRequest{
		Scopes:  []github.Scope{github.ScopeRepo},
		NoteURL: &OAuthAppURL,
	}

	*request.Note, err = authTokenNote()
	if err != nil {
		return
	}

	auth, res, err := client.Authorizations.Create(context.Background(), &request)
	if err != nil {
		return
	}

	if res.StatusCode == 201 {
		token = *auth.Token
	} else {
		errInfo, e := ioutil.ReadAll(res.Body)
		if e == nil {
			err = fmt.Errorf("%s", string(errInfo))
		} else {
			err = e
		}
	}

	return
}

func isToken(password string) bool {
	// api.PrepareRequest = func(req *http.Request) {
	// 	req.Header.Set("Authorization", "token "+password)
	// }

	// res, _ := api.Get("user")
	// if res != nil && res.StatusCode == 200 {
	// 	return true
	// }
	return false
}

func authTokenNote() (string, error) {
	n := os.Getenv("USER")

	if n == "" {
		n = os.Getenv("USERNAME")
	}

	if n == "" {
		whoami := exec.Command("whoami")
		whoamiOut, err := whoami.Output()
		if err != nil {
			return "", err
		}
		n = strings.TrimSpace(string(whoamiOut))
	}

	h, err := os.Hostname()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s for %s@%s", version.AppName, n, h), nil
}
