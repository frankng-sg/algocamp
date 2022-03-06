// Definition for singly-linked list.
var ListNode = /** @class */ (function () {
    function ListNode(val, next) {
        this.val = val === undefined ? 0 : val;
        this.next = next === undefined ? null : next;
    }
    return ListNode;
}());
function detectCycle(head) {
    if (head.next === null)
        return null;
    var slowPtr = head;
    var fastPtr = head;
    while (slowPtr !== null && fastPtr !== null && fastPtr.next !== null) {
        fastPtr = fastPtr.next.next;
        slowPtr = slowPtr.next;
        if (slowPtr === fastPtr)
            break;
    }
    if (slowPtr !== fastPtr)
        return null;
    slowPtr = head;
    while (slowPtr !== fastPtr) {
        slowPtr = slowPtr.next;
        fastPtr = fastPtr.next;
    }
    return slowPtr;
}
var n = new ListNode(1);
console.log(detectCycle(n));
