package main

import (
	"fmt"
	"math"
)

func lowerInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type MinStack struct {
	minVal []int
	last   int
	stack  []int
}

func Constructor() MinStack {
	return MinStack{
		minVal: []int{math.MaxInt},
		last:   0,
	}
}

func (this *MinStack) Push(val int) {
	if len(this.stack) > this.last {
		this.stack[this.last] = val
	} else {
		this.stack = append(this.stack, val)
		this.minVal = append(this.minVal, val)
	}
	if this.last == 0 {
		this.minVal[0] = val
	} else {
		this.minVal[this.last] = lowerInt(this.minVal[this.last-1], val)
	}
	this.last++
}

func (this *MinStack) Pop() {
	this.last--
}

func (this *MinStack) Top() int {
	return this.stack[this.last-1]
}

func (this *MinStack) GetMin() int {
	return this.minVal[this.last-1]
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
