package main

import (
	"fmt"
	"strings"
	"testing"
)

func reverse(head *EduLinkedListNode) *EduLinkedListNode {
	if head == nil || head.next == nil {
		return head
	}
	var prev, curr, next *EduLinkedListNode
	curr = head
	for curr != nil {
		next = curr.next
		curr.next = prev
		prev = curr
		curr = next
	}
	return prev
}

func Test_reverse(t *testing.T) {
	input := [][]int{
		{1, 2, 3, 4, 5},
		{1, 2, 3, 4, 5, 6},
		{3, 2, 1},
		{10},
		{1, 2},
	}

	for i := 0; i < len(input); i++ {
		inputLinkedList := EduLinkedList{}
		inputLinkedList.CreateLinkedList(input[i])
		fmt.Printf("%d. \tInput linked list: ", i+1)
		DisplayLinkedListWithForwardArrow(inputLinkedList.head)
		fmt.Printf("\n\tReversed linked list: ")
		DisplayLinkedListWithForwardArrow(reverse(inputLinkedList.head))
		fmt.Printf("\n%s\n", strings.Repeat("-", 100))
	}

}
