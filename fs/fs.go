// +build go1.16

// fs.FS is only available from go 1.16 onwards

package fs

import (
	actualFs "io/fs"
	"os"
)

// FS We use a type alias to make it easier for the pre-go1.16
// code to fullfil this local type.
type FS = actualFs.FS

// RealFS complies with the fs.FS interface.
// We use this rather than os.DirFS as DirFS has no concept of
// what the current working directory is, whereas if we're a simple
// passthru to os.Open then working dir is automagically taken care of.
type RealFS struct{}

// Open complies with the fs.FS interface
func (dir RealFS) Open(name string) (actualFs.File, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	return f, nil
}

var (
	// ReadDir is an alias to the real implementation. Once the need for backwards compat goes away, so can this.
	ReadDir = actualFs.ReadDir
	// Stat is an alias to the real implementation. Once the need for backwards compat goes away, so can this.
	Stat = actualFs.Stat
	// ReadFile is an alias to the real implementation. Once the need for backwards compat goes away, so can this.
	ReadFile = actualFs.ReadFile
)
