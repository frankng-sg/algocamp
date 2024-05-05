package main

type EduLinkedListNode struct {
	data int
	next *EduLinkedListNode
}

func NewLinkedListNode(data int, next *EduLinkedListNode) *EduLinkedListNode {
	node := new(EduLinkedListNode)
	node.data = data
	node.next = next
	return node
}

func InitLinkedListNode(data int) *EduLinkedListNode {
	node := new(EduLinkedListNode)
	node.data = data
	node.next = nil
	return node
}

func removeNthLastNode(head *EduLinkedListNode, n int) *EduLinkedListNode {
	left, right := head, head
	for i := 0; i < n; i++ {
		right = right.next
	}
	if right == nil {
		return head.next
	}
	for right.next != nil {
		left = left.next
		right = right.next
	}
	left.next = left.next.next
	return head
}
