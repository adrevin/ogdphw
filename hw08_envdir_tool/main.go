package main

import (
	"flag"
	"fmt"
)

func main() {
	flag.Parse()
	flagArgs := flag.Args()
	path := flagArgs[0]
	command := flagArgs[1]
	arg1 := flagArgs[2]
	arg2 := flagArgs[3]
	fmt.Printf("%s %s %s %s\n", path, command, arg1, arg2)

	env, err := ReadDir(path)
	fmt.Printf("%v %s\n", env, err)
}
