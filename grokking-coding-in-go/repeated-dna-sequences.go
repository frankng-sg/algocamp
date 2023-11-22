package main

import (
	"crypto/rand"
	"fmt"
	"math"
	"math/big"
	"sync"
	"time"
)

func repeatedDNASequence(dna string, k int) (out []string) {
	startTime := time.Now()
	if len(dna) < k {
		return
	}

	m := make(map[string]int)
	for i := len(dna); i >= k; i-- {
		m[dna[i-k:i]] += 1
	}
	for key, val := range m {
		if val > 1 {
			out = append(out, key)
		}
	}
	fmt.Printf("use map[string]int: %v\n", time.Since(startTime))
	return
}

func rndStr(n int, charset string) string {

	b := make([]byte, n)
	charsetLen := big.NewInt(int64(len(charset)))
	for i := range b {
		index, err := rand.Int(rand.Reader, charsetLen)
		if err != nil {
			panic(err)
		}
		b[i] = charset[index.Int64()]
	}
	return string(b)
}

func findRepeatedSequences(s string, k int) (out []string) {
	startTime := time.Now()
	n := len(s)

	if n < k {
		return
	}

	mapping := map[rune]int{'A': 1, 'C': 2, 'G': 3, 'T': 4}

	a := 4

	numbers := make([]int, n)
	for i, ch := range s {
		numbers[i] = mapping[ch]
	}

	hashValue := 0

	hashSet := make(map[int]bool)
	var output []string

	for i := 0; i < n-k+1; i++ {

		if i == 0 {
			for j := 0; j < k; j++ {
				hashValue += numbers[j] * int(math.Pow(float64(a), float64(k-j-1)))
			}
		} else {
			previousHashValue := hashValue
			hashValue = ((previousHashValue - numbers[i-1]*int(math.Pow(float64(a), float64(k-1)))) * a) + numbers[i+k-1]
		}

		if _, ok := hashSet[hashValue]; ok {
			output = append(output, s[i:i+k])
		}

		hashSet[hashValue] = true
	}
	fmt.Printf("rolling hash: %v\n", time.Since(startTime))
	return output
}

func main() {
	// fmt.Println(repeatedDNASequence("AAAAACCCCCAAAAACCCCCC", 8))
	str := rndStr(100000000, "ACGT")
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		repeatedDNASequence(str, 10)
	}()
	go func() {
		defer wg.Done()
		findRepeatedSequences(str, 10)
	}()
	wg.Wait()
}
