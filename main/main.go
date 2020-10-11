package main

import (
	"fmt"
	"net/http"
	"studygo"
)

// IndexHandler HTTPリクエストを処理する
func IndexHandler(
	w http.ResponseWriter,
	r *http.Request) {

	fmt.Fprint(w, "hello world")
}

func main() {
	// hello world
	fmt.Println(studygo.ECHO)
	fmt.Println("this is main")

	// HttpServer
	http.HandleFunc("/", IndexHandler)
	http.ListenAndServe(":3000", nil)
}
