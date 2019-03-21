// +build mage

package main

import "github.com/magefile/mage/sh"

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