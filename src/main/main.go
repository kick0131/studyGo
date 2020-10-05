package main

import (
	"fmt"
	"net/http"
	"os"
)

func IndexHandler(
	w http.ResponseWriter,
	r *http.Request) {

	fmt.Fprint(w, "hello world")
}

func main() {
	// hello world
	//fmt.Println(gosample.Message)

	// HttpServer
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println(port)
	http.HandleFunc("/", IndexHandler)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
