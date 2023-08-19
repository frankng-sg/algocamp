package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_countBits(t *testing.T) {
	tests := []struct {
		n    int
		want []int
	}{
		{
			n:    5,
			want: []int{0, 1, 1, 2, 1, 2},
		},
		{
			n:    2,
			want: []int{0, 1, 1},
		},
		{
			n:    0,
			want: []int{0},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("n=%d", tt.n), func(t *testing.T) {
			if got := countBits(tt.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countBits() = %v, want %v", got, tt.want)
			}
		})
	}
}
