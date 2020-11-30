// テストファイルはGoDoc出力対象外
//
// テストに使用するパッケージ
// https://github.com/stretchr/testify
package gosample

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// お試しテスト
func TestHelloWorld(t *testing.T) {
	assert := assert.New(t)

	// 期待値
	expectResult := "ABCDE"

	// モックで動きを変える
	mock := new(MockGoSample)
	mock.On("SubFuncHello").Return(expectResult)

	// テスト実行
	// 本来の動作は上、モックに置き換えたものが下
	// result := SampleFunctionForTest(&SimpleStruct{})
	result := SampleFunctionForTest(mock)

	// 結果確認
	assert.Equal(expectResult, result)
}

func Example() {
	fmt.Println("hello")
	// Output: SubFuncHello
}

func TestGetAllASCII(t *testing.T) {
	assert := assert.New(t)
	result := GetAllASCII()
	fmt.Println(result)
	assert.NotEmpty(result)
}
