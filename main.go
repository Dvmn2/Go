package main

import (
	"fmt"
	"strconv"
)

// test
func main() {
	var a = "hiiii"
	var b = 2
	var c int = 5

	a = "Hi"

	f := ", 2 + 2 = "

	fmt.Println(a + f + strconv.Itoa((b + c)))

}
