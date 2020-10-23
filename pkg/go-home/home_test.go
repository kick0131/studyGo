package gohome

import "testing"

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
