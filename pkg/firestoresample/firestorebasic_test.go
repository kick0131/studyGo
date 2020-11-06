package firestoresample

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

type CredentialData struct {
	ProjectName string `json:"project_id"`
}

// JSONPATH はGoogleSDKのサービスアカウントファイルです
const JSONPATH = "../../data/adminsdk.json"

// TestWriteFs はFirestoreの書き込みテストを行います
func TestWriteFs(t *testing.T) {
	// 環境変数を設定し、接続先をFirestoreエミュレータに変更
	os.Setenv("FIRESTORE_EMULATOR_HOST", "localhost:8080")

	// 認証情報ファイルの絶対パス
	jsonpath, _ := filepath.Abs(JSONPATH)
	fmt.Println(jsonpath)

	// Firestoreオブジェクトを生成し、書き込み処理を実行
	client := BuildClientWithCredential(jsonpath)
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
