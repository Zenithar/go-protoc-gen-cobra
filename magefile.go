// +build mage

package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/magefile/mage/sh"
)

var curDir = func() string {
	name, _ := os.Getwd()
	return name
}()

// Calculate file paths
var toolsBinDir = normalizePath(path.Join(curDir, "tools", "bin"))

func init() {
	time.Local = time.UTC

	// Add local bin in PATH
	err := os.Setenv("PATH", fmt.Sprintf("%s:%s", toolsBinDir, os.Getenv("PATH")))
	if err != nil {
		panic(err)
	}
}

func BuildExample() error {
	return sh.Run(
		"protoc",
		"-I", "/usr/local/include",
		"-I", "./example",
		"--go_out=plugins=grpc:./example",
		"./example/example.proto",
	)
}

func BuildCobra() error {
	return sh.Run(
		"protoc",
		"-I", "/usr/local/include",
		"-I", "./example",
		"--cobra_out=:./example",
		"./example/example.proto",
	)
}

// ------------------------------------------------------------------------------------

// normalizePath turns a path into an absolute path and removes symlinks
func normalizePath(name string) string {
	absPath := mustStr(filepath.Abs(name))
	return absPath
}

func mustStr(r string, err error) string {
	if err != nil {
		panic(err)
	}
	return r
}
