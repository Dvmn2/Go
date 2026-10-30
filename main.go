package main

import "fmt"

func fib(c int) []int {
	a := 0
	b := 1
	c -= 2
	array := []int{0, 1}
	for i := 0; i < c; i++ {
		t := a
		a = b
		b = a + t
		array = append(array, b)
	}
	return array
}

func fib_rec(array []int, c int) []int {
	if c == 0 {
		return array
	}
	if len(array) == 0 {
		array = []int{0, 1}
		c -= 2
	} else if len(array) == 1 {
		array = append(array, 1)
		c--
	}
	array = append(array, array[len(array)-1]+array[len(array)-2])
	c--
	return fib_rec(array, c)
}

func main() {
	c := 0
	fmt.Println("Enter count:")
	fmt.Scan(&c)
	array := fib(c)
	fmt.Printf("array: %v\n", array)
	array = []int{}
	array = fib_rec(array, c)
	fmt.Printf("array: %v\n", array)
}
