// Definition for singly-linked list.
class ListNode {
    val: number;
    next: ListNode | null;
    constructor(val?: number, next?: ListNode | null) {
        this.val = val === undefined ? 0 : val;
        this.next = next === undefined ? null : next;
    }
}

function detectCycle(head: ListNode | null): ListNode | null {
    if (head === null || head.next === null) return null;

    let slowPtr = head;
    let fastPtr = head;

    while (slowPtr !== null && fastPtr !== null && fastPtr.next !== null) {
        fastPtr = fastPtr.next.next;
        slowPtr = slowPtr.next;
        if (slowPtr === fastPtr) break;
    }

    if (slowPtr !== fastPtr) return null;

    slowPtr = head;
    while (slowPtr !== fastPtr) {
        slowPtr = slowPtr.next;
        fastPtr = fastPtr.next;
    }
    return slowPtr;
}

let n = new ListNode(1);
console.log(detectCycle(n));
