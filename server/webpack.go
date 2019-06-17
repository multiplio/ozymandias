package server

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
)

type manifest struct {
	Files map[string]string `json:"files"`
}

func readManifest(buildPath string) (*manifest, error) {
	assetsManifestPath := path.Join(buildPath, "asset-manifest.json")

	if _, err := os.Stat(assetsManifestPath); os.IsNotExist(err) {
		return new(manifest), nil
	}

	content, err := ioutil.ReadFile(assetsManifestPath)
	if err != nil {
		return nil, err
	}

	var manifest manifest
	if err = json.Unmarshal(content, &manifest); err != nil {
		return nil, err
	}

	return &manifest, nil
}

func (m *manifest) files() map[string]bool {
	files := make(map[string]bool)

	for _, file := range m.Files {
		files[file] = true
	}

	return files
}
