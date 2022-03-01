# Definition for singly-linked list
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution:
    def addTwoNumbers(self, l1: ListNode, l2: ListNode) -> ListNode:
        carry = 0
        # Record the address of the head of the sum
        head = tail = ListNode(0)  # dummy nodes
        while True:
            if l1 is None and l2 is None and carry is 0:
                return head.next

            if l1 is None:
                val1 = 0
            else:
                val1 = l1.val
                l1 = l1.next

            if l2 is None:
                val2 = 0
            else:
                val2 = l2.val
                l2 = l2.next

            digits_sum = val1 + val2 + carry
            new_digit = ListNode(digits_sum % 10)
            carry = digits_sum // 10
            tail.next = new_digit
            tail = new_digit
