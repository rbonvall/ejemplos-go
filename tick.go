package main

import "fmt"
import "time"

func main() {
    t := time.Tick(3 * time.Second)

    for {
        <-t
        fmt.Println("Hola")
    }
}
