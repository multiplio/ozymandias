package api

import (
	"encoding/json"
	"net/http"

	// "github.com/google/go-github/v25/github"
	// "golang.org/x/oauth2"

	"github.com/multiplio/ozymandias/auth"
)

// Handler returns http.Handler for API endpoint
func Handler(user auth.UserReader) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		if !user.AssertAuthed(res, req) {
			return
		}

		res.Header().Set("Content-Type", "application/json")

		body, err := json.Marshal(map[string]interface{}{
			"data": "Hello, world",
		})
		if err != nil {
			res.WriteHeader(500)
			return
		}

		res.WriteHeader(200)
		res.Write(body)
	}
}

// // connect to github api
// ctx := context.Background()
// ts := oauth2.StaticTokenSource(
// 	&oauth2.Token{AccessToken: accessToken},
// )
// tc := oauth2.NewClient(ctx, ts)

// client := github.NewClient(tc)

// // list all organizations
// orgs, _, err := client.Organizations.List(ctx, "", nil)
// if err != nil {
// 	panic(err)
// }
// fmt.Print("%v", orgs)

// // list all repositories
// repos, _, err := client.Repositories.List(ctx, "", nil)
// if err != nil {
// 	panic(err)
// }
// fmt.Print("%v", repos)
