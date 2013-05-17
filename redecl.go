package main
import "fmt"

func main() {
    a, b := f()
    fmt.Println(a)
    fmt.Println(b)
    fmt.Println()

    a, c := f()
    fmt.Println(a)
    fmt.Println(c)
    fmt.Println()
}

func f() (int, int) {
    return 2, 3
}
