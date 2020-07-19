package main

import (
	"io"
	"os"
	"strconv"
)

type duplicator struct {
	src    src
	dst    dst
	offset offset
	limit  limit
}

func newDuplicator(src, dst, filename string, o offset, l limit) (duplicator, error) {
	s, err := newSrc(src)
	if err != nil {
		return duplicator{}, err
	}
	d, err := newDst(s, dst, filename)
	if err != nil {
		return duplicator{}, err
	}
	return duplicator{
		src:    s,
		dst:    d,
		offset: o,
		limit:  l,
	}, nil
}

func (dp duplicator) duplicate() error {
	srcFile, err := os.Open(dp.src.path.string())
	if err != nil {
		return err
	}
	defer srcFile.Close()

	var (
		offset = dp.offset.int()
		limit  = dp.limit.int() + offset
	)
	for i := offset; i < limit; i++ {
		dp.dst.addFilenameSuffixInt(i)
		dstFile, err := os.Create(dp.dst.path.string())
		if err != nil {
			return err
		}
		defer dstFile.Close()

		if _, err = io.Copy(dstFile, srcFile); err != nil {
			return err
		}
		if _, err = srcFile.Seek(0, io.SeekStart); err != nil {
			return err
		}
	}
	return nil
}

type offset int

func newOffset(str string) (offset, error) {
	n, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}
	if n < 0 {
		n = 0
	}
	return offset(n), nil
}

func (o offset) int() int {
	return int(o)
}

type limit int

func newLimit(str string) (limit, error) {
	n, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}
	if n > upperLimit {
		return 0, errOverUpperLimit
	}
	if n < 1 {
		n = 1
	}
	return limit(n), nil
}

func (l limit) int() int {
	return int(l)
}
