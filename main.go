package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world!")
}

func main() {
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8080", nil)
}
