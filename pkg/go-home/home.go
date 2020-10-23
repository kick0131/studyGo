package gohome

import (
	"fmt"

	"github.com/common-nighthawk/go-figure"
)

// GoHome おうちに帰りたい
func GoHome() {
	figure.NewFigure("Go Home!", "basic", true).Scroll(20000, 300, "right")
}

// GoHome2 おうちに帰りたい2
func GoHome2() {
	fmt.Println("aaa")
}
