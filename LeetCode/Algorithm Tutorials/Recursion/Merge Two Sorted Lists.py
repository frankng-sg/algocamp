# Definition for singly-linked list.
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution:
    def mergeTwoLists(self, nodeList1, nodeList2):

        # tail of the merged list should be known throughout recursion (not None)
        def recur_merge(tail, nodeList1, nodeList2):
            if nodeList1 is None:
                tail.next = nodeList2
                return
            if nodeList2 is None:
                tail.next = nodeList1
                return

            if nodeList1.val < nodeList2.val:
                tail.next = nodeList1
                return recur_merge(nodeList1, nodeList1.next, nodeList2)
            else:
                tail.next = nodeList2
                return recur_merge(nodeList2, nodeList1, nodeList2.next)

        dummy_tail = ListNode(0)  # declare a dummy node
        recur_merge(dummy_tail, nodeList1, nodeList2)

        return dummy_tail.next
