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
	"log"
	"unicode/utf8"

	"github.com/common-nighthawk/go-figure"
)

// Foo は変数
var Foo = "fooooo"

// BAR は定数
const BAR = "BAAAAA"

// SimpleIf はシンプルなインタフェース定義です
type SimpleIf interface {
	SubFuncHello() string
}

// SimpleStruct はシンプルな構造体です
type SimpleStruct struct{}

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
func SampleFunctionForTest(inter SimpleIf) string {
	msg := "SampleFunctionForTest"
	fmt.Println(msg)
	return inter.SubFuncHello()
}

// SubFuncHello モック化対象インタフェース実装
func (SimpleStruct) SubFuncHello() string {
	msg := "SubFuncHello"
	fmt.Println(msg)
	return msg
}

// GetAllASCII はGo言語で全アスキーコードの出力を行い、文字列を返す実験メソッドです
func GetAllASCII() string {
	// 文字アクセスにはコードポイントを扱うrune型を使う
	loop := 127 // 0x7FまでがASCIIコード
	var value []rune
	for loop > 0 {
		value = append(value, rune(loop))
		loop--
	}
	fmt.Println(utf8.RuneCountInString(string(value)))
	return string(value)
}

// DoPrintLog logの出力先変更のお試しコード
func DoPrintLog() {
	msg := "DoPrintLog"
	log.Println(msg)
}
