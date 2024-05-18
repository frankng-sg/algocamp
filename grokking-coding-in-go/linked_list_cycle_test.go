package main

func detectCycle(head *EduLinkedListNode) bool {
	if head == nil {
		return false
	}
	slow, fast := head, head.next
	for fast != nil && fast.next != nil {
		if slow == fast {
			return true
		}
		slow = slow.next
		fast = fast.next.next
	}
	return false
}
