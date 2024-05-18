package main

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
