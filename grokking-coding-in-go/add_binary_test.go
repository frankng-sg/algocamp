package main

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func reverseBytes(a []byte) []byte {
	i, j := 0, len(a)-1
	for i < j {
		a[i], a[j] = a[j], a[i]
		i++
		j--
	}
	return a
}

func addBinary(str1 string, str2 string) string {
	var result bytes.Buffer
	var carry byte
	i, j := len(str1)-1, len(str2)-1
	for i >= 0 || j >= 0 {
		sum := carry
		if i >= 0 {
			sum += str1[i] - '0'
		}
		if j >= 0 {
			sum += str2[j] - '0'
		}
		carry = 0
		if sum > 1 {
			sum, carry = sum-2, 1
			carry = 1
		}
		result.WriteByte(sum + '0')
		i, j = i-1, j-1
	}
	if carry > 0 {
		result.WriteByte('1')
	}
	return string(reverseBytes(result.Bytes()))
}

func TestAddBinary(t *testing.T) {
	str1 := []string{"1001", "1", "11"}
	str2 := []string{"1", "1", "11"}
	expect := []string{"1010", "10", "110"}
	for i := 0; i < len(expect); i++ {
		assert.Equal(t, expect[i], addBinary(str1[i], str2[i]))
	}
}
