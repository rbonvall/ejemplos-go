// go run rest.go
// open localhost:8787/api/number in browser

package main

import (
    "code.google.com/p/gorest"
    "net/http"
)


func main() {
    gorest.RegisterService(new(TestService))
    http.Handle("/", gorest.Handle())
    http.ListenAndServe(":8787", nil)
}

type Person struct {
    Given_name string
    Family_name string
}

type TestService struct {

    // Service level configuration
    gorest.RestService `root:"/api/"
                        consumes:"application/json"
                        produces:"application/json"`

    // End-point level configuration
    getNumber gorest.EndPoint `method:"GET"
                               path:"/number/"
                               output:"int"`
    getNumbers gorest.EndPoint `method:"GET"
                                path:"/numbers/"
                                output:"[]int"`
    getPerson gorest.EndPoint `method:"GET"
                               path:"/person/"
                               output:"Person"`
}

func (s TestService) GetNumber() int {
    return 1
}

func (s TestService) GetNumbers() []int {
    return []int{1, 2, 3, 4, 5}
}

func (s TestService) GetPerson() Person {
    return Person{"Perico", "Los Palotes"}
}
