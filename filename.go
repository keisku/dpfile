package main

import (
	"path/filepath"
	"strings"
)

type filename struct {
	basename  string
	extention string
}

func newFilename(v string) (filename, error) {
	if v == "" {
		return filename{}, nil
	}
	base := filepath.Base(v)
	dirbase := filepath.Base(filepath.Dir(v))
	if base == dirbase {
		return filename{}, errSrcIsNotFilePath
	}
	ext := filepath.Ext(v)
	return filename{
		basename:  strings.Replace(base, ext, "", 1),
		extention: ext,
	}, nil
}

func (fn *filename) merge(src src) {
	if fn.hasName() && fn.equalName(src.filename.name()) {
		fn.addSuffix("_duplicated")
	}
	if !fn.hasName() {
		fn.applyName(src.filename.name())
		fn.addSuffix("_duplicated")
	}
	fn.applyExt(src.filename.ext())
}

func newEmptyFilename() filename {
	return filename{}
}

func (fn filename) string() string {
	return fn.name() + fn.ext()
}

func (fn filename) name() string {
	return fn.basename
}

func (fn filename) ext() string {
	return fn.extention
}

func (fn *filename) addSuffix(s string) {
	fn.basename += s
}

func (fn *filename) applyName(n string) {
	fn.basename += n
}

func (fn *filename) applyExt(ext string) {
	fn.extention += ext
}
func (fn filename) hasName() bool {
	return fn.basename != ""
}

func (fn filename) hasExt() bool {
	return fn.extention != ""
}

func (fn filename) equalName(n string) bool {
	return fn.basename == n
}
