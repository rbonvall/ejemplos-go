package main

import "time"
import "fmt"

func sleepsort(input []int) (output []int) {
    output = make([]int, len(input))

    c := make(chan int)
    for _, x := range input {
        go func(x int) {
            time.Sleep(time.Duration(x) * time.Millisecond)
            c <- x
        }(x)
    }

    i := 0
    for _, _ = range input {
        output[i] = <-c
        i++
    }
    return
}

func main() {
    var a, b []int

    a = []int{6, 1, 4, 3, 5, 2}

    fmt.Println(a)
    b = sleepsort(a)
    fmt.Println(b)

}
