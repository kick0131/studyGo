package gohome

import (
	"github.com/common-nighthawk/go-figure"
)

// GoHome おうちに帰りたい
func GoHome() {
	figure.NewFigure("Go Home!", "basic", true).Scroll(20000, 300, "right")
}
