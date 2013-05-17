package main
import "fmt"

func main() {
	var n int
	fmt.Printf("Gimme an integer: ")
	fmt.Scanf("%d", &n)

	// Doesn't work.
	//if 1 < n < 10 {
	//	fmt.Println("It is between 1 and 10")
	//} else {
	//	fmt.Println("It is not between 1 and 10")
	//}

	if 1 < n && n < 10 {
		fmt.Println("It is between 1 and 10")
	} else {
		fmt.Println("It is not between 1 and 10")
	}
}
