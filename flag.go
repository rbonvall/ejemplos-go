package main

import (
    "fmt"
    "flag"
)

func main() {
    n := flag.Int("test", 1000, "A test value.")
    flag.Parse()
    fmt.Printf("%v tiene el valor %v\n", "test", *n)
    flag.PrintDefaults()

}
