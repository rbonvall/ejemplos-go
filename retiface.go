package main
import "fmt"

func f(x int) interface{} {
	if x % 2 == 0 {
		return "hola"
	}
	return 3.14159
}

func main() {
	a := f(1)
	b := f(2)
	fmt.Println("a:", a)
	fmt.Println("b:", b)

}
