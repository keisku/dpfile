package main

type path struct {
	dir      dir
	filename filename
}

func newPath(d dir, fn filename) path {
	return path{dir: d, filename: fn}
}

func (p path) string() string {
	return p.dir.string() + "/" + p.filename.string()
}

func (p path) applyFilename(fn filename) path {
	p.filename = fn
	return p
}

type src struct{ path }

func newSrc(path string) (src, error) {
	dir, err := newDir(path)
	if err != nil {
		return src{}, err
	}
	fn, err := newFilename(path)
	if err != nil {
		return src{}, err
	}
	if !fn.hasName() {
		return src{}, errNotFoundFileName
	}
	if !fn.hasExt() {
		return src{}, errNotFoundExt
	}
	return src{newPath(dir, fn)}, nil
}

type dst struct{ path }

func newDst(src src, destination, filename string) (dst, error) {
	dir, err := newDir(destination)
	if err != nil {
		return dst{}, err
	}
	fn, err := newFilename(filename)
	if err != nil {
		return dst{}, err
	}
	fn.merge(src)
	return dst{newPath(dir, fn)}, nil
}
