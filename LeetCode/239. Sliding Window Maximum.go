package main

import "fmt"

func maxSlidingWindow(nums []int, k int) []int {
	output := []int{}
	q, qlen := []int{}, 0
	l, r, n := 0, 0, len(nums)
	for r < n {
		for qlen > 0 && nums[q[qlen-1]] < nums[r] {
			q, qlen = q[:qlen-1], qlen-1
		}
		q, qlen = append(q, r), qlen+1
		if l > q[0] {
			q, qlen = q[1:], qlen-1
		}
		if r+1 >= k {
			output = append(output, nums[q[0]])
			l++
		}
		r++
	}
	return output
}

func main() {
	fmt.Println(maxSlidingWindow([]int{1, 2, 3, 3, 2, 1, 1, 4}, 3)) // Output: [3, 3, 3, 3, 2, 4]
}
