package main

import (
	"fmt"
	"os"

	vm "github.com/vcokltfre/gosemby/vm"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s <file>\n", os.Args[0])
		os.Exit(1)
	}

	fileName := os.Args[1]

	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	vm.ExecBytecode(data, 256, 256)
}
