// Definition for singly-linked list.
class ListNode {
    val: number;
    next: ListNode | null;
    constructor(val?: number, next?: ListNode | null) {
        this.val = val === undefined ? 0 : val;
        this.next = next === undefined ? null : next;
    }
}

function hasCycle(head: ListNode | null): boolean {
    let slowPtr = head;
    let fastPtr = head;

    while (slowPtr !== null && fastPtr !== null && fastPtr.next !== null) {
        fastPtr = fastPtr.next.next;
        slowPtr = slowPtr.next;
        if (slowPtr === fastPtr) return true;
    }
    return false;
}
