package fs1

import (
	"testing"
	"testing/fstest"

	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
)

var (
	testFilesystem = fstest.MapFS{
		"main.go": &fstest.MapFile{
			Data: []byte(`
package main

import "foo/bar"

func main() {
	bar.PrintSomething()
}
`),
		},
		"_pkg/src/foo/bar/bar.go": &fstest.MapFile{
			Data: []byte(`
package bar

import (
	"fmt"
)

func PrintSomething() {
	fmt.Println("I am printing something!")
}
`),
		},
	}
)

func TestFunctionCall(t *testing.T) {
	i := interp.New(interp.Options{
		GoPath:     "./_pkg",
		Filesystem: testFilesystem,
	})
	if err := i.Use(stdlib.Symbols); err != nil {
		t.Fatal(err)
	}

	_, err := i.EvalPath(`main.go`)
	if err != nil {
		t.Fatal(err)
	}
}
