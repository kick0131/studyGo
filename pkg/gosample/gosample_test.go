// テストファイルはGoDoc出力対象外
package gosample

import (
	"fmt"
	"testing"
)

// お試しテスト
func TestHelloWorld(t *testing.T) {
	SampleFunctionForTest()
}

func Example() {
	fmt.Println("hello")
	// Output: hello
}
