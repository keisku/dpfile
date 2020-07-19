package main

import "errors"

var (
	errSrcIsNotFilePath = errors.New("You should set the file path to --src, -s")
	errDstIsNotDir      = errors.New("You should set the directory to --dst, -d")
	errOverUpperLimit   = errors.New("You should set the limit within 10000")
	errBrokenFile       = errors.New("The file you designated is broken")
)

const (
	upperLimit = 10000
)
