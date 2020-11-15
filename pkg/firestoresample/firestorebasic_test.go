package firestoresample

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
	"time"

	"google.golang.org/api/iterator"
)

// TestDataSampleA はFirestore動作確認用のマッピング構造体です
type TestDataSampleA struct {
	A string `firestore:"first"`
	B string `firestore:"last"`
	C int64  `firestore:"born"`
}

// CredentialData 認証情報のjsonファイルマッピング先となる構造体です（未使用）
type CredentialData struct {
	ProjectName string `json:"project_id"`
}

// Firestore投入情報（データパターン１）
var docData = NoticeInfo{
	subCollection: []NoticeInfoSub{
		{
			collectionName: "notice_jp",
			CreateAt:       time.Now(),
			Title:          "モーニングコール",
			Data:           "おはよう!!",
			Dispflag:       true,
		},
		{
			collectionName: "notice_en",
			CreateAt:       time.Now(),
			Title:          "Morning call",
			Data:           "good morning!!",
			Dispflag:       true,
		},
		{
			collectionName: "notice_vn",
			CreateAt:       time.Now(),
			Title:          "cuộc gọi buổi sáng",
			Data:           "Buổi sáng tốt lành!!",
			Dispflag:       true,
		},
	},
}

// Firestore投入情報（データパターン２）
var docData2 = NoticeInfo{
	subCollection: []NoticeInfoSub{
		{
			collectionName: "notice_jp",
			CreateAt:       time.Now(),
			Title:          "自分の国",
			Data:           "日本",
			Dispflag:       true,
		},
		{
			collectionName: "notice_en",
			CreateAt:       time.Now(),
			Title:          "My country",
			Data:           "America",
			Dispflag:       true,
		},
		{
			collectionName: "notice_vn",
			CreateAt:       time.Now(),
			Title:          "Đất nước của tôi",
			Data:           "Việt Nam",
			Dispflag:       true,
		},
	},
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
	// 有効にした場合、接続先をFirestoreエミュレータに変更する
	os.Setenv("FIRESTORE_EMULATOR_HOST", "localhost:8080")

	// 認証情報ファイルの絶対パス
	jsonpath, _ := filepath.Abs(JSONPATH)
	fmt.Println(jsonpath)

	// Firestoreオブジェクトを生成し、書き込み処理を実行
	client = BuildClientWithCredential(jsonpath)

}

func shutdown() {
	client.fb.Close()
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

// TestReadFsWithStruct はFirestoreから構造体へのマッピングを行います
func TestReadFsWithStruct(t *testing.T) {
	fmt.Println("=== CheckPoint ReadFsWithStruct 1")

	var docData TestDataSampleA

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
		doc.DataTo(&docData)
		fmt.Println(docData)
		fmt.Printf("docData.A :%s\n", docData.A)
		fmt.Printf("docData.B :%s\n", docData.B)
		fmt.Printf("docData.C :%#v\n", docData.C)
	}
}

// TestWriteFsWithStruct は構造体からFirestoreへのマッピングを行います
func TestWriteFsWithStruct(t *testing.T) {
	fmt.Println("=== CheckPoint WriteFsWithStruct 1")

	// 定義されないフィールドはomitemptyフラグにより生成されない
	createtime, _ := time.Parse("2006-01-02 15:04:05 MST", "2020-11-11 11:22:33 JST")
	// loc, _ := time.LoadLocation("Asia/Tokyo")
	// createtime := time.Now()
	docData := NoticeInfoSub{
		CreateAt: createtime,
		Title:    "Morning call",
		Data:     "good morning!!",
		Dispflag: false,
	}

	// 書き込み(ドキュメントIDは自動生成)
	_, err := client.fb.Collection("noticeInfo").NewDoc().Create(client.ctx, docData)
	if err != nil {
		t.Fatalf("document Create error\n")
	}
}

// TestWriteFsSubCollection はサブコレクションのデータ書き込みを行います
func TestWriteFsSubCollection(t *testing.T) {
	fmt.Println("=== CheckPoint WriteFsSubCollection 1")

	// サブコレクション定義は外部で実装可

	// サブコレクション生成
	err := docData.Set(client, "noticeInfo")
	err = docData2.Set(client, "noticeInfo")
	if err != nil {
		t.Fatalf("sub-collection Create error\n")
	}
}

// TestReadFsSubCollection はサブコレクションのデータ取得を行います
func TestReadFsSubCollection(t *testing.T) {
	fmt.Println("=== CheckPoint ReadFsSubCollection 1")

	// サブコレクションのデータ取得
	noticeInfo := NoticeInfo{}
	err := noticeInfo.GetSubcollectionDoc(client, "notice_jp")
	if err != nil {
		t.Fatalf("GetSubcollectionDoc error\n")
	}

	// 取得内容確認
	for _, item := range noticeInfo.subCollection {
		fmt.Printf("title : %s\n", item.Title)
		fmt.Printf("data  : %s\n", item.Data)
		fmt.Printf("flag  : %v\n", item.Dispflag)
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
