/*
Package main is my package
*/
package main

import (
	"fmt"
	"net/http"
	"os"
	"studygo/pkg/gosample"
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

	// gosample.HelloWorld()
	// gohome.GoHome()
	// gosample.RamdamIntn()
	fmt.Println(gosample.YearToZodiac(1982))
	// tourofgo.Tofgo01()

	// fmt.Println(quote.Opt())
}
