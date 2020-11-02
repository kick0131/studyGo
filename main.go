/*
Package main is my package
*/
package main

import (
	"fmt"
	"net/http"
	"os"
	"studygo/pkg/gosample"
	"studygo/posgredata"
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
	year := 2005
	fmt.Printf("year %d zodiac is %s\n", year, gosample.YearToZodiac(year))
	fmt.Printf("random : %s\n", gosample.RamdamDispZodiac())

	// PostgresQL
	posgredata.SampleQuery()

	// fmt.Println(quote.Opt())
}
