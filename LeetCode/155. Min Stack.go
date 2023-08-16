package main

import (
	"fmt"
)

func lowerInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type MinStack struct {
	n      int
	minVal []int
	stack  []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	if this.n == 0 {
		this.minVal = append(this.minVal, val)
	} else {
		this.minVal = append(this.minVal, lowerInt(this.minVal[this.n-1], val))
	}
	this.n++
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:this.n-1]
	this.minVal = this.minVal[:this.n-1]
	this.n--
}

func (this *MinStack) Top() int {
	return this.stack[this.n-1]
}

func (this *MinStack) GetMin() int {
	return this.minVal[this.n-1]
}

func main() {
	minStack := &MinStack{}
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	fmt.Println(minStack.GetMin()) // return -3
	minStack.Pop()
	fmt.Println(minStack.Top())    // return 0
	fmt.Println(minStack.GetMin()) // return -2
}
