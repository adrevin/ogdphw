package main

import (
	"golang.org/x/example/stringutil"
)

const hello = "Hello, OTUS!"

func main() {
	print(stringutil.Reverse(hello))
}
