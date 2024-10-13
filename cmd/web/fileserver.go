package main

import (
	"net/http"
	"strings"
)

type protectedFileSystem struct {
	fs http.FileSystem
}

func (pfs protectedFileSystem) Open(path string) (http.File, error) {
	f, err := pfs.fs.Open(path)
	if err != nil {
		return nil, err
	}

	s, err := f.Stat()
	if s.IsDir() {
		index := strings.TrimSuffix(path, "/") + "/index.html"
		if _, err := pfs.fs.Open(index); err != nil {
			return nil, err
		}
	}

	return f, nil
}
