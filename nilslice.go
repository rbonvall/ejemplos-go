package main

import "fmt"

type T struct {
	x, y int
	s string
	z []int
}


func e() []int  { return nil }
func f() int    { return nil }
func g() *int   { return nil }
func h() string { return nil }
func i() T      { return nil }

func main() {
	fmt.Println(f())
	fmt.Println(g())
	fmt.Println(h())
	fmt.Println(i())
}
