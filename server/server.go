package server

import (
	"net/http"
)

// Handler returns http.Handler for server endpoint
func Handler(buildPath string) http.Handler {
	return http.FileServer(http.Dir(buildPath))
}
