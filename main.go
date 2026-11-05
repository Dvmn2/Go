package main

import (
	"fmt"
	"iter"
	"strings"
)

type Vector[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next, prev *element[T]
	val        T
}

func (a *Vector[T]) Pull(val T) {
	e := &element[T]{val: val}
	if a.head == nil {
		a.head = e
		a.tail = e
	} else {
		e.prev = a.tail
		a.tail.next = e
		a.tail = e
	}
}

func (a *Vector[T]) Put(val T) {
	e := &element[T]{val: val}
	if a.head == nil {
		a.head = e
		a.tail = e
	} else {
		e.next = a.head
		a.head.prev = e
		a.head = e
	}
}

func (a *Vector[T]) String() string {
	var s strings.Builder
	for i := a.head; i != nil; i = i.next {
		s.WriteString(fmt.Sprint(i.val))
		s.WriteString(" ")
	}
	return s.String()
}

func (vec *Vector[T]) Elements() iter.Seq[T] {
	return func(yield func(T) bool) {
		for e := vec.head; e != nil; e = e.next {
			if !yield(e.val) {
				return
			}
		}
	}
}

func Squaring() iter.Seq[int] {
	return func(yield func(int) bool) {
		a := 1
		for {
			if !yield(a) {
				return
			}
			a *= 2
		}
	}
}

func Squaring2() iter.Seq2[int, int] {
	return func(yield func(int, int) bool) {
		a := 1
		for i := 0; ; i++ {
			if !yield(i, a) {
				return
			}
			a *= 2
		}
	}
}

func main() {
	vector := &Vector[int]{}
	vector.Pull(1)
	vector.Pull(2)
	vector.Put(0)

	for n := range vector.Elements() {
		fmt.Println(n)
	}

	fmt.Println("---")

	for n := range Squaring() {
		if n >= 10 {
			break
		}
		fmt.Println(n)
	}

	fmt.Println("---")

	for n, v := range Squaring2() {
		if n >= 10 {
			break
		}
		fmt.Println(v)
	}
}
