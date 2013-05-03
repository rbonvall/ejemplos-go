package main

import "fmt"

func main() {
	s := "select %v from students"
	t := fmt.Sprintf(s, "\b\b\b\b\b\b\b delete *")

	fmt.Println("Before:")
	for _, c := range([]byte(s)) {
		fmt.Printf("%c ", c)
	}
	fmt.Printf("\n\n")

	fmt.Println("After:")
	for _, c := range([]byte(t)) {
		fmt.Printf("%c ", c)
	}
	fmt.Printf("\n\n")

	// No, it doesn't.
}
