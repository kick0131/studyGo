package firestoresample

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"google.golang.org/api/iterator"
)

// TestDataSampleA はFirestore動作確認用のマッピング構造体です
type TestDataSampleA struct {
	A string `firestore:"first,omitempty"`
	B string `firestore:"last,omitempty"`
	C string `firestore:"born,omitempty"`
}

// CredentialData 認証情報のjsonファイルマッピング先となる構造体です（未使用）
type CredentialData struct {
	ProjectName string `json:"project_id"`
}

// JSONPATH はGoogleSDKのサービスアカウントファイルです
const JSONPATH = "../../data/adminsdk.json"

var client *FireBaseClient

// TestMain はFiresotreにテスト用テーブルを用意し、終了後にテーブルを削除します
func TestMain(m *testing.M) {
	// 準備処理
	setup()

	// m.Runでテストの終了を待つ
	exitCode := m.Run()

	// 終了処理
	shutdown()

	os.Exit(exitCode)
}

func setup() {
	// 環境変数を設定し、接続先をFirestoreエミュレータに変更
	os.Setenv("FIRESTORE_EMULATOR_HOST", "localhost:8080")

	// 認証情報ファイルの絶対パス
	jsonpath, _ := filepath.Abs(JSONPATH)
	fmt.Println(jsonpath)

	// Firestoreオブジェクトを生成し、書き込み処理を実行
	client = BuildClientWithCredential(jsonpath)

}

func shutdown() {

}

// TestWriteFs はFirestoreの書き込みテストを行います
func TestWriteFs(t *testing.T) {
	fmt.Println("=== CheckPoint Write 1")

	// コレクションusersにドキュメント追加
	_, _, err := client.fb.Collection("users").Add(client.ctx, map[string]interface{}{
		"first": "Ada",
		"last":  "Lovelace",
		"born":  1815,
	})
	if err != nil {
		t.Log(err)
		fmt.Println(err)
	}
}

// TestReadFs はFirestoreの読み込みテストを行います
func TestReadFs(t *testing.T) {
	fmt.Println("=== CheckPoint Read 1")

	// コレクション内のドキュメント一覧を取得
	iter := client.fb.Collection("users").Documents(client.ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			fmt.Println("=== itereator.Done")
			break
		}
		if err != nil {
			t.Fatalf("document read error\n")
			// return err
		}
		fmt.Println(doc.Data())
	}
}

// jsonファイルの内容を構造体にマッピング
func filetojson(filepath string) *CredentialData {
	raw, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var credentialData CredentialData

	json.Unmarshal(raw, &credentialData)
	return &credentialData
}
