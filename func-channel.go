package main

import (
	"fmt"
)

type T func (int) int

func gr(ch chan T) {
	for {
		f := <-ch
		fmt.Println(f(5))
	}
}

func main() {
	ch := make(chan T)
	go gr(ch)
	ch <- func (x int) int { return x * x }
	ch <- func (x int) int { return x + 2 }
	ch <- func (x int) int { return 60 }
}
