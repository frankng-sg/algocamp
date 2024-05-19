package main

func getMiddleNode(head *EduLinkedListNode) *EduLinkedListNode {
	if head == nil {
		return nil
	}
	fast, slow := head, head
	for fast != nil && fast.next != nil {
		fast = fast.next.next
		slow = slow.next
	}
	return slow
}
