package main

import "fmt"

func main() {

	//i := 1
	//for i <= 3 {
	//	fmt.Println(i)
	//	i++
	//}
	//
	//for j := 0; j < 3; j++ {
	//	fmt.Println(j)
	//}
	//
	//for i := range 3 {
	//	fmt.Println("range", i)
	//}
	//
	//for {
	//	fmt.Println("loop", i)
	//	break
	//}
	//
	//for n := range 6 {
	//	if n%2 == 0 {
	//		continue
	//	}
	//	fmt.Println(n)
	//}

	// wtf
	k := 1
	for i := 0; i < 2; i++ {
		fmt.Println(k)
		k := "k"
		fmt.Println(k)
	}
	fmt.Println(k)
}
