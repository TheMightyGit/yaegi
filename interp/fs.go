// +build go1.16

// fs.FS is only available from go 1.16

package interp

import (
	"io/fs"
	"os"
)

type FS = fs.FS

// realFS complies with the fs.FS interface.
// We use this rather than os.DirFS as DirFS has no concept of
// what the current working directory is, whereas if we're a simple
// passthru to os.Open then working dir is automagically taken care of.
type realFS struct{}

func (dir realFS) Open(name string) (fs.File, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	return f, nil
}
