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
	result := SampleFunctionForTest()
	t.Error("this is Error code")
	fmt.Println(result)
	assert.Equal(result, "hoge", "they should be equal")
}

func TestHelloWorld2(t *testing.T) {
	assert := assert.New(t)
	result := SampleFunctionForTest()
	t.Error("this is Success code")
	fmt.Println(result)
	assert.Equal(result, "huga", "they should be equal")
}

func Example() {
	fmt.Println("hello")
	// Output: hello
}

func TestGetAllASCII(t *testing.T) {
	assert := assert.New(t)
	result := GetAllASCII()
	fmt.Println(result)
	assert.NotEmpty(result)
}
