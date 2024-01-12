package main

import (
	"fmt"
	"net/http"

	"github.com/nuxzero/sample/cal"
)

func main() {
	fmt.Println("Hello World")
	result := cal.Add(1, 2)
	fmt.Println(result)
}

func restAPI() {
	http.NewServeMux()
}

func handleRequest() {

}
