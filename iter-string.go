package main
import "fmt"

func main() {
	a := "hola mundo ¢rüëł"
	for i, r := range a {
		fmt.Printf("%3d '%v'\n", i, string(r))
	}
}
