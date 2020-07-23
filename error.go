package main

import "errors"

var (
	errSrcIsNotFilePath = errors.New("You should set the file path to --src, -s")
	errNotFoundExt      = errors.New("The extension of file is not found")
	errNotFoundFileName = errors.New("The name of file is not found")
	errOverUpperLimit   = errors.New("You should set the limit within 10000")
)

const (
	upperLimit = 10000
)
