// テストファイルはGoDoc出力対象外
//
// テストに使用するパッケージ
// https://github.com/stretchr/testify
package gosample

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// お試しテスト
func TestHelloWorld(t *testing.T) {
	assert := assert.New(t)

	// 期待値
	expectResult := "ABCDE"

	// モックで動きを変える
	mock := new(MockGoSample)
	mock.On("SubFuncHello").Return(expectResult)

	// テスト実行
	// 本来の動作は上、モックに置き換えたものが下
	// result := SampleFunctionForTest(&SimpleStruct{})
	result := SampleFunctionForTest(mock)

	// 結果確認
	assert.Equal(expectResult, result)
}

func Example() {
	fmt.Println("SubFuncHello")
	// Output: SubFuncHello
}

// なにかしら値が入っていればOKとし、中身の評価はスキップ
func TestGetAllASCII(t *testing.T) {
	assert := assert.New(t)
	result := GetAllASCII()
	fmt.Println(result)
	assert.NotEmpty(result)
}

// ログの出力内容を確認する
func TestPrintLog(t *testing.T) {
	assert := assert.New(t)

	// 期待値
	expectResult := "DoPrintLog"

	// 標準出力先をバッファに変更
	var buf bytes.Buffer
	log.SetOutput(&buf)
	// デフォルトだと日付が出力されてしまうので、フラグに0を設定する
	// 	const (
	//     Ldate         = 1 << iota     // ローカルタイムゾーンでの日付: 2009/01/23
	//     Ltime                         // ローカルタイムゾーンでの時刻: 01:23:23
	//     Lmicroseconds                 // マイクロ秒: 01:23:23.123123.  Ltime も設定されていることを仮定しています。
	//     Llongfile                     // フルファイル名と行番号: /a/b/c/d.go:23
	//     Lshortfile                    // ファイル名と行番号: d.go:23. Llongfile を上書きします。
	//     LUTC                          // Ldate または Ltime が設定されている場合，ローカルタイムゾーンではなく UTC を使います。
	//     LstdFlags     = Ldate | Ltime // 標準ロガーの初期値
	// )
	defaultFlags := log.Flags()
	log.SetFlags(0)
	defer func() {
		log.SetOutput(os.Stderr)
		log.SetFlags(defaultFlags)
	}()

	// テスト
	DoPrintLog()
	actual := strings.TrimRight(buf.String(), "\n")

	// 結果確認
	assert.Equal(expectResult, actual)
}

// ExforSample 拡張for文のサンプル
func TestExforSample(t *testing.T) {
	type StudentInfo struct {
		name string
		age  int
	}

	students := []StudentInfo{
		{"sum", 10},
		{"bob", 18},
		{"mary", 24},
	}

	// rangeの1要素ずつループさせる構文(foreachと同等)
	for idx, item := range students {
		log.Println(fmt.Sprintln("[", idx, "] name:", item.name, "age:", item.age))
	}

}

// 文字列セパレート
// go test -v -run "TestSepalate" -count=1 ./pkg/gosample/.
func TestSepalateSample(t *testing.T) {
	assert := assert.New(t)

	// 変換元
	emails := []string{
		"user1@sample.com",
		"hoge@fuga.co.jp",
		"ABCDE@F.G.H.I",
	}

	// 変換元文字列のカンマ結合
	emailStr := ""
	for i, item := range emails {
		if i == 0 {
			emailStr = item
		} else {
			emailStr += ("," + item)
		}
	}
	log.Println(emailStr)

	// カンマ結合した文字列の分割
	actual := []string{}
	slice := strings.Split(emailStr, ",")
	for _, item := range slice {
		actual = append(actual, strings.TrimSpace(item))
	}

	// 元の文字配列と比較
	assert.Equal(emails, actual)
}
