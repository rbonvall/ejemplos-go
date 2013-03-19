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

type TestService struct {

    // Service level configuration
    gorest.RestService `root:"/api/"
                        consumes:"application/json"
                        produces:"application/json"`

    // End-point level configuration
    number gorest.EndPoint `method:"GET"
                            path:"/number/"
                            output:"int"`
}

func (s TestService) Number() int {
    return 1
}
