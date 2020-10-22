/*
Package main is my package
*/
package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/common-nighthawk/go-figure"

	"rsc.io/quote"
)

// HTTPServer サーバ
func HTTPServer() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println(port)
	http.HandleFunc("/", IndexHandler)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}

// IndexHandler リクエストハンドラ
func IndexHandler(
	w http.ResponseWriter,
	r *http.Request) {

	fmt.Fprint(w, "hello world")
}

func main() {
	// hello world
	// fmt.Println(gosample.Message)

	myFigure := figure.NewFigure("Hello World", "", true)
	myFigure.Print()

	fmt.Println(quote.Opt())
}
