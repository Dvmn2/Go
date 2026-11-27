package main

import (
	"fmt"
	"strings"
)

func reverse(slice []string) []string {
	for i := range slice[0 : len(slice)/2] {
		slice[i], slice[len(slice)-i-1] = slice[len(slice)-i-1], slice[i]
	}
	return slice
}

func remove(slice []string, id int) []string {
	slice = append(slice[:id], slice[id+1:]...)
	return slice
}

func wordsCount(str string) map[string]int {
	count := make(map[string]int)
	for word := range strings.SplitSeq(str, " ") {
		count[word]++
	}

	return count
}

func groupCount(count map[string]int) map[int][]string {
	group := make(map[int][]string)
	for word, c := range count {
		group[c] = append(group[c], word)
	}
	return group
}

func f3() {
	req := string("123456")
	s := make([]string, 6)
	s = strings.Split(req, "")
	fmt.Println(s, len(s), cap(s))
	s = reverse(s)
	fmt.Println(s, len(s), cap(s))
	s = remove(s, 2)
	fmt.Println(s, len(s), cap(s))

	str := "test1 test2 test3 test1 test3"
	count := wordsCount(str)
	fmt.Println(count)
	for word, c := range count {
		if c > 1 {
			fmt.Println(word)
		}
	}
	group := groupCount(count)
	fmt.Printf("groupCount(count): %v\n", group)
}

func main() {
	f3()
}
