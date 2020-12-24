/*
Package main is my package
*/
package main

import (
	"fmt"
	"studygo/pkg/gosample"
	"studygo/pkg/httpserver"
	// "studygo/pkg/posgredata"
)

// パッケージを使うものだけ記載する
// 単に動かすだけならパッケージ内のテストコードから呼び出す
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
	// posgredata.SampleQuery()

	// fmt.Println(quote.Opt())
	httpserver.HTTPServer()
}
