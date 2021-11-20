package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func main() {
	input := "Hello, OTUS!"
	fmt.Println(reverse(input))
}

func reverse(input string) string {
	return stringutil.Reverse(input)
}
