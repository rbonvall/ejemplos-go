package main

import "fmt"

type Cosa struct {
    n int
    s string
}


func main() {

    var a, b *Cosa
    a = &Cosa{1, "uno"}
    b = &Cosa{2, "dos"}

    fmt.Printf("a: %v, b: %v\n", a, b)
    a, b = b, a
    fmt.Printf("a: %v, b: %v\n", a, b)
    a, b = a, nil
    fmt.Printf("a: %v, b: %v\n", a, b)
}
