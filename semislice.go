package main

import "fmt"

func main() {
	a := []int{3, 1, 4, 1, 5, 9}

	for _, x := range a[2:] {
		fmt.Println(x)
	}
}
