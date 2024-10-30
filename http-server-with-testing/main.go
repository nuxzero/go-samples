package main

import (
	"net/http"

	"github.com/nuxzero/go-samples/foo"
)

func main() {
	http.HandleFunc("/", foo.Greeting)
	http.HandleFunc("/goodbye", foo.SayGoodbye)
	http.HandleFunc("/talk", foo.Talk)

	http.ListenAndServe(":8080", nil)
}
