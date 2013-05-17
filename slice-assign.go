package main
import "fmt"

func main() {
	var s, t []int

	s = append(s, 10)
	s = append(s, 20)
	s = append(s, 30)

	t = s

	fmt.Println(s, t)
}
