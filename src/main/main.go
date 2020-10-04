package main

import (
	"fmt"
	"gosample"
	"net/http"
)

func IndexHandler(
	w http.ResponseWriter,
	r *http.Request) {

	fmt.Fprint(w, "hello world")
}

func main() {
	// hello world
	fmt.Println(gosample.Message)

	// HttpServer
	http.HandleFunc("/", IndexHandler)
	http.ListenAndServe(":3000", nil)
}
