package server

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type errorHandler struct {
	err error
}

func (e errorHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	log.Println(e.err)

	res.WriteHeader(http.StatusInternalServerError)
	fmt.Fprint(res, e.err)
}

type reactHandler struct {
	Path  string
	Files map[string]bool
}

func (d reactHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var filename string

	if req.URL.Path == "/" {
		filename = "/index.html"
	} else if _, present := d.Files[req.URL.Path]; present {
		filename = filepath.FromSlash(path.Clean(req.URL.Path))
	} else {
		filename = "/index.html"
	}
	fullName := filepath.Join(d.Path, filename)
	fmt.Println(fullName)

	f, err := os.Open(fullName)
	if err != nil {
		fmt.Fprint(res, mapDirOpenError(err, fullName))
		return
	}

	io.Copy(res, f)
	return
}

// mapDirOpenError maps the provided non-nil error from opening name
// to a possibly better non-nil error. In particular, it turns OS-specific errors
// about opening files in non-directories into os.ErrNotExist.
func mapDirOpenError(originalErr error, name string) error {
	if os.IsNotExist(originalErr) || os.IsPermission(originalErr) {
		return originalErr
	}

	parts := strings.Split(name, string(filepath.Separator))
	for i := range parts {
		if parts[i] == "" {
			continue
		}

		fi, err := os.Stat(strings.Join(parts[:i+1], string(filepath.Separator)))
		if err != nil {
			return originalErr
		}
		if !fi.IsDir() {
			return os.ErrNotExist
		}
	}

	return originalErr
}

// Handler returns http.Handler for ui files
func Handler(buildPath string) http.Handler {
	manifest, err := readManifest(buildPath)
	if err != nil {
		return errorHandler{err}
	}

	return http.Handler(reactHandler{buildPath, manifest.files()})
}
