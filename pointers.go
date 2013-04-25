package main

import (
	"fmt"
)

type T struct {
	s *string
}

func main() {
	var a, b int
	var p, q *int

	a = 15
	b = 20

	p = new(int)
	*p = 25

	q = &b

	fmt.Printf("a: %d\tb: %d\t*p: %d\t*q: %d\n", a, b, *p, *q)

	var t *T
	var r string

	r = "hola amigos"
	t = new(T)
	t.s = new(string)
	*t.s = r

	fmt.Printf("t: %v\t*t.s: %v\n", t, *t.s)
}
