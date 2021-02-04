package fcmsubscribe

import (
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

// HTTPServer サーバ
func HTTPServer() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println(port)
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/subscribe", subscribeHandler)
	http.HandleFunc("/unsubscribe", unsubscribeHandler)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}

// IndexHandler リクエストハンドラ
func IndexHandler(
	w http.ResponseWriter,
	r *http.Request) {

	// fmt.Println("method:", r.Method) //リクエストを取得するメソッド
	if r.Method == "GET" {
		// テンプレートの内容を出力する
		t, err := template.ParseFiles("./pkg/fcmsubscribe/template.html")
		if err != nil {
			log.Fatalf("template error: %v", err)
		}
		t.Execute(w, nil)
	} else {
		fmt.Fprintf(w, "HTTP Getメソッドを使ってください")
	}
}

// getJSONFromContentBody 共通チェックとリクエストボディのボディ長を取得します
func getJSONFromContentBody(
	w http.ResponseWriter,
	r *http.Request) (map[string]interface{}, error) {
	//Validate request
	if r.Method != "POST" {
		fmt.Println("Not POST")
		w.WriteHeader(http.StatusBadRequest)
		return nil, errors.New("Method Check")
	}

	if r.Header.Get("Content-Type") != "application/json" {
		fmt.Println("Not application/json ")
		w.WriteHeader(http.StatusBadRequest)
		return nil, errors.New("Content-Type")
	}
	//To allocate slice for request body
	length, err := strconv.Atoi(r.Header.Get("Content-Length"))
	if err != nil {
		fmt.Println("Not content-length ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return nil, errors.New("Content-Length")
	}

	//Read body data to parse json
	body := make([]byte, length)
	length, err = r.Body.Read(body)
	if err != nil && err != io.EOF {
		w.WriteHeader(http.StatusInternalServerError)
		return nil, errors.New("get Body")
	}

	//parse json
	var jsonBody map[string]interface{}
	err = json.Unmarshal(body[:length], &jsonBody)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return nil, errors.New("parse json")
	}

	return jsonBody, nil
}

// subscribeHandler リクエストのトークンIDをトピックに登録します
func subscribeHandler(
	w http.ResponseWriter,
	r *http.Request) {

	// 共通チェック+HTTPボディ取得
	jsonBody, err := getJSONFromContentBody(w, r)
	if err != nil {
		return
	}

	// fmt.Printf("%v\n", jsonBody)
	// fmt.Printf("TokenID:%v\n", jsonBody["TokenID"])
	// fmt.Printf("Topic:%v\n", jsonBody["Topic"])

	// トピックと登録するデバイストークンを生成
	registrationTokens := []string{
		jsonBody["TokenID"].(string),
	}
	topic := jsonBody["Topic"].(string)

	// トピック登録
	jsonpath, _ := filepath.Abs(JSONPATH)
	message := BuildClientWithCredential(jsonpath)
	result, err := message.SubscribeToTopic(registrationTokens, topic)
	if err != nil {
		fmt.Fprintf(w, "%v\n", err)
		return
	}
	fmt.Fprintf(w, "%v\n", result)

	w.WriteHeader(http.StatusOK)
}

// unsubscribeHandler リクエストのトークンIDをトピックから解除します
func unsubscribeHandler(
	w http.ResponseWriter,
	r *http.Request) {

	// 共通チェック+HTTPボディ取得
	jsonBody, err := getJSONFromContentBody(w, r)
	if err != nil {
		return
	}

	// fmt.Printf("%v\n", jsonBody)
	// fmt.Printf("TokenID:%v\n", jsonBody["TokenID"])
	// fmt.Printf("Topic:%v\n", jsonBody["Topic"])

	// トピックと登録するデバイストークンを生成
	registrationTokens := []string{
		jsonBody["TokenID"].(string),
	}
	topic := jsonBody["Topic"].(string)

	// トピック開放
	jsonpath, _ := filepath.Abs(JSONPATH)
	message := BuildClientWithCredential(jsonpath)
	result, err := message.UnsubscribeFromTopic(registrationTokens, topic)
	if err != nil {
		fmt.Fprintf(w, "Error %v\n", err)
		return
	}
	fmt.Fprintf(w, "Successfly %v\n", result)

	w.WriteHeader(http.StatusOK)
}
