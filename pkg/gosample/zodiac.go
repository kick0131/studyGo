package gosample

import (
	"fmt"
	"math/rand"
	"time"
)

// 干支のインデックス定義
// iotaはconst内で使用する演算子0から自動で値を割り当ててくれる
const (
	// 干支：子
	MOUSE = iota
	// 干支：丑
	OX
	// 干支：寅
	TIGER
	// 干支：兎
	HARE
	// 干支：竜
	DRAGON
	// 干支：巳
	SERPENT
	// 干支：午
	HORSE
	// 干支：未
	SHEEP
	// 干支：申
	MONKEY
	// 干支：酉
	ROOSTAR
	// 干支：戌
	DOG
	// 干支：亥
	BOAR
)

// Zodiac は干支を表します
type Zodiac struct {
	// 干支に対応するインデックス値
	Index int
	// 干支名
	Name string
}

// changeYearToZodiacIndex は西暦を何番目の干支かを表すインデックスに変換します
func changeYearToZodiacIndex(year int) int {
	// 最初の干支が0になるように補正をかける
	return (year + 8) % 12
}

// YearToZodiac は西暦を干支の文字列に変換して返します
func YearToZodiac(year int) string {
	return IndexToZodiac(changeYearToZodiacIndex(year))
}

// IndexToZodiac はインデックス化された干支を対応する干支の文字列に変換して返します
func IndexToZodiac(index int) string {
	sercher := []Zodiac{
		{Index: MOUSE, Name: "MOUSE"},
		{Index: OX, Name: "OX"},
		{Index: TIGER, Name: "TIGER"},
		{Index: HARE, Name: "HARE"},
		{Index: DRAGON, Name: "DRAGON"},
		{Index: SERPENT, Name: "SERPENT"},
		{Index: HORSE, Name: "HORSE"},
		{Index: SHEEP, Name: "SHEEP"},
		{Index: MONKEY, Name: "MONKEY"},
		{Index: ROOSTAR, Name: "ROOSTAR"},
		{Index: DOG, Name: "DOG"},
		{Index: BOAR, Name: "BOAR"},
	}
	for i, v := range sercher {
		fmt.Printf("index:%d value:%s\n", i, v.Name)
		if index == v.Index {
			return v.Name
		}
	}
	return "Unknown"
}

// RamdamDispZodiac はいずれかの干支をランダムで出力します
func RamdamDispZodiac() {
	rand.Seed(time.Now().UnixNano())

}

// RamdamIntn はランダムな数字を標準出力します
func RamdamIntn() {
	rand.Seed(time.Now().UnixNano())
	for i := 1; i < 10; i++ {
		fmt.Println(rand.Intn(10))
	}
}
