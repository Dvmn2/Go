package main

import (
	"fmt"
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

func main() {
	vector := &Vector[int]{}
	vector.Pull(1)
	vector.Pull(2)
	vector.Put(0)
	fmt.Println(vector) // 0 1 2
}
