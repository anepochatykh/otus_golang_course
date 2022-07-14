package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func main() {
	line := "Hello, OTUS!"
	reversedLine := stringutil.Reverse(line)
	fmt.Println(reversedLine)
}
