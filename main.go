package main

import (
	"fmt"
	"math/rand"
)

func l1(n int, m int) [][]int {
	array := make([][]int, n)
	for i := range n {
		array[i] = make([]int, m)
	}

	for i := range n {
		fmt.Printf("Enter %d row: ", i)
		for j := range m {
			fmt.Scan(&array[i][j])
		}
	}

	return array
}

func l2(n int, m int) [][]int {
	array := make([][]int, n)
	for i := range n {
		array[i] = make([]int, m)
	}

	for i := range n {
		for j := range m {
			array[i][j] = rand.Intn(50)*2 + (i+j)%2
		}
	}

	return array
}

func main() {
	n, m := 0, 0
	fmt.Println("Enter n and m:")
	fmt.Scan(&n)
	fmt.Scan(&m)

	array := l1(n, m)
	fmt.Printf("array: %v\n", array)
	array = l2(n, m)
	fmt.Printf("array: %v\n", array)
}
