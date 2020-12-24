package httpserver

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
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

	fmt.Println("method:", r.Method) //リクエストを取得するメソッド
	if r.Method == "GET" {
		// テンプレートの内容を出力する
		t, err := template.ParseFiles("pkg/httpserver/template.html")
		if err != nil {
			log.Fatalf("template error: %v", err)
		}
		t.Execute(w, nil)
	} else {
		fmt.Fprintf(w, "POSTはダメよ")
	}
}
