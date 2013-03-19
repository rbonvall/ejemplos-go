package main
import "fmt"

type Cosa struct {
    n int
    s string
}

func main() {
    var list []Cosa

    c1 := Cosa{1, "uno"}
    list = append(list, c1)
    
    c2 := Cosa{2, "dos"}
    list = append(list, c2)

    for i, x := range list {
        fmt.Printf("%v: %v\n", i, x)
    }
}
