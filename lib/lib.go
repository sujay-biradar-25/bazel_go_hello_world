package mylib

import (
	"fmt"

	"github.com/sujay-biradar-25/bazel_go_hello_world/internal_deps"
)

func MyFunction() {
	fmt.Println("Hello from MyFunction!")
	internal_deps.InternalFunction()
}
