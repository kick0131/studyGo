package gosample

import "testing"

// 評価しないで単に実行するだけ
func TestSimpleGorutine01(t *testing.T) {
	target := SimpleStruct{}
	target.SimpleGorutine01()
}
func TestSimpleGorutine02(t *testing.T) {
	target := SimpleStruct{}
	target.SimpleGorutine02()
}

// func TestSimpleGorutine03(t *testing.T) {
// 	target := SimpleStruct{}
// 	target.SimpleGorutine03()
// }
