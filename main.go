package main

import (
	"fmt"
	"math"
)

const s = "const"

// test
func main() {
	fmt.Println(s)

	const a = 50000
	const b = 3e20 / a
	fmt.Println(int64(b))
	fmt.Println(math.Sin(a))

}
