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

	// == pkg/gohome
	// gohome.GoHome()

	// == pkg/gosample
	year := 2005
	fmt.Printf("year %d zodiac is %s\n", year, gosample.YearToZodiac(year))
	fmt.Printf("random : %s\n", gosample.RamdamDispZodiac())

	// == pkg/httpserver
	httpserver.HTTPServer()

	// == pkg/posgredata
	// posgredata.SampleQuery()
	// fmt.Println(quote.Opt())
}
