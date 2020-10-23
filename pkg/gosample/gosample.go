/*

Package gosample テスト練習パッケージ

あいうえお
かきくけこ

	アイウエオ
	カキクケコ

さしすせそ

	コメントは type, variable, constant, function, packageに記載可能

*/
package gosample

import (
	"fmt"

	"github.com/common-nighthawk/go-figure"
)

// Foo は変数
var Foo = "fooooo"

// BAR は定数
const BAR = "BAAAAA"

// HelloWorld figureモジュールを使った文字列出力
//
// Info
//
// 複数行あった場合の表示確認
// 		サンプル
func HelloWorld() {
	myFigure := figure.NewFigure("Hello World", "", true)
	myFigure.Print()
}

// SampleFunctionForTest go test(Example)のための関数
func SampleFunctionForTest() string {
	fmt.Println("hoge")
	return "huga"
}
