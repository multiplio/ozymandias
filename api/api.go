package api

import (
	"encoding/json"
	"net/http"

	"github.com/multiplio/ozymandias/auth"
)

// Handler returns http.Handler for API endpoint
func Handler(user UserReader) http.HandlerFunc {
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
