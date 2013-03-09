package main
import "fmt"

type Cosa struct {
    int
    float64
    string
}

func main() {
    var c Cosa
    c.int = 2
    c.float64 = 3.14
    c.string = "Dábale arroz a la zorra el abad."

    fmt.Println(c)
}

