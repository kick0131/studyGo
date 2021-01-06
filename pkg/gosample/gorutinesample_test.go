package gosample

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// 評価しないで単に実行するだけ
func TestSimpleGorutine01(t *testing.T) {
	target := SimpleStruct{}
	target.SimpleGorutine01()
}
func TestSimpleGorutine02(t *testing.T) {
	target := SimpleStruct{}
	target.SimpleGorutine02()
}

// sync.waitGroupを使ったゴルーチンの待機処理サンプル
// go test -v -count=1 -run "TestSleepMethod" .\pkg\gosample\.
//
// 課題
// ・タイムアウトが取れない
// ・異常系が取れていない
//
// 基本的な流れ
// ------------------
// WaitGroup変数の宣言
// Addでカウントアップ
// Doneでカウントダウン
// Waitで待機
func TestSleepMethod1(t *testing.T) {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		// 処理結果を待たない場合、ゴルーチンを使わないケースより応答が早くなる
		go func() {
			defer wg.Done()
			time.Sleep(time.Second * 1)
		}()
	}
	wg.Wait()
	fmt.Println("end")
}

// gorutine自体をgorutineで呼ぶ
// →並列実行となる(10並列x10並列)
func TestSleepMethod2(t *testing.T) {
	target := SimpleStruct{}
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			target.SleepMethodGorutine()
		}()
	}
	wg.Wait()
	fmt.Println("end")
}
