package main

import (
	"os"
	"path/filepath"
)

type dir string

func newDir(d string) (dir, error) {
	if _, err := os.Stat(d); err != nil {
		return "", err
	}
	return dir(filepath.Dir(d)), nil
}

func (d dir) string() string {
	return string(d)
}
