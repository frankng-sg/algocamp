package main

import "fmt"

type Queue []int

func (this *Queue) Enqueue(v int) {
	*this = append(*this, v)
}

func (this *Queue) Dequeue() int {
	v := (*this)[0]
	*this = (*this)[1:]
	return v
}

func (this *Queue) Pop() {
	*this = (*this)[:len(*this)-1]
}

func (this *Queue) Top() int {
	return (*this)[len(*this)-1]
}

func (this *Queue) Front() int {
	return (*this)[0]
}

func (this *Queue) Empty() bool {
	return len(*this) == 0
}

func maxSlidingWindow(nums []int, k int) []int {
	output := []int{}
	q := Queue{}
	l, r, n := 0, 0, len(nums)
	for r < n {
		for !q.Empty() && nums[q.Top()] < nums[r] {
			q.Pop()
		}
		q.Enqueue(r)
		if l > q.Front() {
			q.Dequeue()
		}
		if r+1 >= k {
			output = append(output, nums[q.Front()])
			l++
		}
		r++
	}
	return output
}

func main() {
	fmt.Println(maxSlidingWindow([]int{1, 2, 3, 3, 2, 1, 1, 4}, 3)) // Output: [3, 3, 3, 3, 2, 4]
}
