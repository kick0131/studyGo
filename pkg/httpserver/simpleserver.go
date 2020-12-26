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
	http.HandleFunc("/setupsi", setupSIHandler)
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

// setupSIHandler リクエストハンドラ
func setupSIHandler(
	w http.ResponseWriter,
	r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatalf("ParseForm error: %v", err)
	}

	// Destination はラジオボタンで選択したデプロイ先環境名です
	Destination := r.Form.Get("dstmode")
	// FrontEndHost はテキストフィールドのフロントエンドホストアドレスです
	FrontEndHost := r.Form.Get("FrontEndHost")
	// BackEndHost はテキストフィールドのバックエンドホストアドレスです
	BackEndHost := r.Form.Get("BackEndHost")
	// Emails はテキストエリアのメールアドレスです（セパレータはカンマ）
	Emails := r.Form.Get("Emails")

	fmt.Printf("Destination    : %s\n", Destination)
	fmt.Printf("FrontEndHost   : %s\n", FrontEndHost)
	fmt.Printf("BackEndHost    : %s\n", BackEndHost)
	fmt.Printf("Emails         : %s\n", Emails)

	fmt.Fprintf(w, fmt.Sprintf("Destination  : %s\r\n", Destination))
	fmt.Fprintf(w, fmt.Sprintf("FrontEndHost : %s\r\n", FrontEndHost))
	fmt.Fprintf(w, fmt.Sprintf("BackEndHost  : %s\r\n", BackEndHost))

}
