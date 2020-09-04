package main

import (
	"go-boilerplate/internals/container"
)

func main() {
	sv, err := container.NewContainer()
	if err != nil {
		panic(err)
	}
	if err := sv.Start(); err != nil {
		panic(err)
	}
}
