// Code generated by github.com/izumin5210/gex. DO NOT EDIT.

// +build tools

package tools

// tool dependencies
import (
	_ "github.com/gobuffalo/packr/packr"
	_ "github.com/golang/protobuf/protoc-gen-go"
	_ "github.com/izumin5210/gex/cmd/gex"
)

// If you want to use tools, please run the following command:
//  go generate ./tools.go
//
//go:generate go build -v -o=./bin/gex github.com/izumin5210/gex/cmd/gex
