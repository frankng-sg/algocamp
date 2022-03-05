var __spreadArray = (this && this.__spreadArray) || function (to, from, pack) {
    if (pack || arguments.length === 2) for (var i = 0, l = from.length, ar; i < l; i++) {
        if (ar || !(i in from)) {
            if (!ar) ar = Array.prototype.slice.call(from, 0, i);
            ar[i] = from[i];
        }
    }
    return to.concat(ar || Array.prototype.slice.call(from));
};
var LinkedListNode = /** @class */ (function () {
    function LinkedListNode(data) {
        this.data = data;
        this.next = null;
    }
    return LinkedListNode;
}());
var MyLinkedList = /** @class */ (function () {
    function MyLinkedList() {
        this.linkedList = null;
        this.length = 0;
    }
    MyLinkedList.prototype.getNodeAtIndex = function (index) {
        var p = 0;
        var node = this.linkedList;
        while (p < index && node !== null) {
            node = node.next;
            p++;
        }
        return node;
    };
    MyLinkedList.prototype.get = function (index) {
        var node = this.getNodeAtIndex(index);
        if (node === null)
            return -1;
        return node.data;
    };
    MyLinkedList.prototype.addAtHead = function (val) {
        var newNode = new LinkedListNode(val);
        newNode.next = this.linkedList;
        this.linkedList = newNode;
        this.length++;
    };
    MyLinkedList.prototype.addAtTail = function (val) {
        var newNode = new LinkedListNode(val);
        if (this.length === 0) {
            this.linkedList = newNode;
        }
        else {
            var tail = this.getNodeAtIndex(this.length - 1);
            tail.next = newNode;
        }
        this.length++;
    };
    MyLinkedList.prototype.addAtIndex = function (index, val) {
        if (index > this.length)
            return;
        if (index == 0) {
            this.addAtHead(val);
        }
        else {
            var nodeBeforeIndex = this.getNodeAtIndex(index - 1);
            var nodeAfterIndex = nodeBeforeIndex.next;
            var newNode = new LinkedListNode(val);
            nodeBeforeIndex.next = newNode;
            newNode.next = nodeAfterIndex;
            this.length++;
        }
    };
    MyLinkedList.prototype.deleteAtIndex = function (index) {
        if (index > this.length)
            return;
        if (index == 0) {
            this.linkedList = this.linkedList.next;
        }
        else {
            var nodeBeforeIndex = this.getNodeAtIndex(index - 1);
            var nodeAfterIndex = nodeBeforeIndex.next.next;
            nodeBeforeIndex.next = nodeAfterIndex;
        }
        this.length--;
    };
    MyLinkedList.prototype.toString = function () {
        var output = [];
        var node = this.linkedList;
        while (node !== null) {
            output.push(node.data);
            node = node.next;
        }
        return output.toString();
    };
    return MyLinkedList;
}());
var obj = new MyLinkedList();
// const cmd = ["MyLinkedList", "addAtHead", "get", "addAtHead", "addAtHead", "deleteAtIndex", "addAtHead", "get", "get", "get", "addAtHead", "deleteAtIndex"];
// const args = [[], [4], [1], [1], [5], [3], [7], [3], [3], [3], [1], [4]];
// const cmd = ["MyLinkedList", "addAtHead", "addAtTail", "addAtIndex", "get", "deleteAtIndex", "get"];
// const args = [[], [1], [3], [1, 2], [1], [1], [1]];
var cmd = ["MyLinkedList", "addAtHead", "deleteAtIndex", "addAtHead", "addAtHead", "addAtHead", "addAtHead", "addAtHead", "addAtTail", "get", "deleteAtIndex", "deleteAtIndex"];
var args = [[], [2], [1], [2], [7], [3], [2], [5], [5], [5], [6], [4]];
for (var i = 1; i < cmd.length; i++) {
    console.log("------------------------------");
    console.log.apply(console, __spreadArray(["Command: ", cmd[i]], args[i], false));
    console.log("Output: ", obj[cmd[i]].apply(obj, args[i]));
    console.log("Linked list: ", obj.toString());
}
