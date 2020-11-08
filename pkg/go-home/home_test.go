package gohome

import (
	"fmt"
	"testing"
)

func TestGoHome(t *testing.T) {
	// GoHome()
	t.Fatal("this method is fatal")
}

func TestGoHome2(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			name: "hoge",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GoHome2()
		})
	}
}

func ExampleGoHome() {
	fmt.Println("Hello")
	// output: Hello
}
