package main

import "fmt"

type T struct {
	n int
	x float32
}

func (t *T) printStuff() {
	fmt.Println("Hola amigazo")
}

func main() {
	var t *T = nil
	t.printStuff()
}

