// Create a directory tree for testing purposes.

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

func create_file(args ...string) {
	ioutil.WriteFile(path.Join(args...), []byte("asdfdfcs\n"), 0644)
}

func main() {
	base, _ := ioutil.TempDir(".", "testfolder")
	//defer os.RemoveAll(base)

	fmt.Println(base)

	os.Mkdir(path.Join(base, "Illapu"), 0755)
	create_file(base, "Illapu", "Lejos del amor.mp3")
	create_file(base, "Illapu", "Sincero positivo.mp3")

	create_file(base, "informe.doc")
	create_file(base, "foto.jpg")
}
