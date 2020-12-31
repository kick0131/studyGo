package gosample

import "fmt"

// SimpleGorutine01 はゴルーチンとチャネルの基本的な使用例です
func (*SimpleStruct) SimpleGorutine01() {

	// 容量0のチャネルを定義
	c := make(chan int)

	// これをやると固まる。チャネルを使えるのはゴルーチン内のみ（？）
	// c <- 10

	// 引数の合計値をチャネルに送信するゴルーチン
	go func(s []int, c chan int) {
		sum := 0
		for _, v := range s {
			sum += v
		}

		// チャネルへ送信
		c <- sum
	}([]int{1, 3, -2}, c)

	// チャネルから受信
	x := <-c
	fmt.Println(x)
}

// SimpleGorutine02 はゴルーチンのサンプルメソッドです
func (*SimpleStruct) SimpleGorutine02() {
	// 容量を指定したチャネル
	ch1 := make(chan int, 1)
	// 組み込みメソッドcapでチャネルの容量を取得
	// 組み込みメソッドlenはチャネル内の変数サイズを取得
	fmt.Println("cap:", cap(ch1), " len:", len(ch1))

	ch2 := make(chan string, 5)
	fmt.Println("cap:", cap(ch2), " len:", len(ch2))

	// 容量0の場合はポップするまでロックするが、容量が1あるのでここでロックしない
	ch1 <- 10
	for _, item := range []string{"first", "second", "third", "forth", "fifth"} {
		ch2 <- item
	}
	fmt.Println("cap:", cap(ch2), " len:", len(ch2))

	fmt.Println("ch1:", <-ch1, " ch2:", <-ch2)
}
