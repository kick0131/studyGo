package gosample

import (
	"github.com/common-nighthawk/go-figure"
)

// HelloWorld figureモジュールを使った文字列出力
func HelloWorld() {
	myFigure := figure.NewFigure("Hello World", "", true)
	myFigure.Print()
}
