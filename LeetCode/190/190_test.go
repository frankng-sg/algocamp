package main

import (
	"fmt"
	"testing"
)

func Test_reverseBits(t *testing.T) {
	tests := []struct {
		num  uint32
		want uint32
	}{
		{
			num:  0,
			want: 0,
		},
		{
			num:  1 << 31,
			want: 1,
		},
		{
			num:  0b01010101010101010101010101010101,
			want: 0b10101010101010101010101010101010,
		},
		{
			num:  0b11110000111100001111000011110000,
			want: 0b00001111000011110000111100001111,
		},
		{
			num:  0b00000000111111110000000011111111,
			want: 0b11111111000000001111111100000000,
		},
		// Add more test cases here
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("num=%d", tt.num), func(t *testing.T) {
			if got := reverseBits(tt.num); got != tt.want {
				t.Errorf("reverseBits() = %v, want %v", got, tt.want)
			}
		})
	}
}
