package main

import (
	"llvm.org/llvm/bindings/go/llvm"
)

func main() {
	_ = llvm.NewModule("foo")
}
