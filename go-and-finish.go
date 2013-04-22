package main

import (
	"time"
	"fmt"
)

func wait(n int) {
	time.Sleep(time.Duration(n) * time.Second)
	fmt.Println("Holi")
}

func main() {

	go wait(10)
	go wait(20)

	var n int
	fmt.Scanf("%d", &n)
}
