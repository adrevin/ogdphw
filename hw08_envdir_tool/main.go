package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	env, err := ReadDir(args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	os.Exit(RunCmd(args[2:], env))
}
