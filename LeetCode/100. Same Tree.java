// Definition for a binary tree node.
 public class TreeNode {
     int val;
     TreeNode left;
     TreeNode right;
     TreeNode(int x) { val = x; }
 }

class Solution {
    public boolean isSameTree(TreeNode p, TreeNode q) {
    	Stack<TreeNode> stack = new Stack<TreeNode>();
        TreeNode node1, node2;
        stack.push(p);
      	stack.push(q);
      	do {
          	node1 = stack.pop();
          	node2 = stack.pop();
          	if (node1 == null && node2 == null) {
             	continue;
            }
          	if (node1 == null || node2 == null) {
              	return false;
            }
          	if (node1.val != node2.val) {
              	return false;
            }

          	stack.push(node1.left);
          	stack.push(node2.left);
          	stack.push(node1.right);
          	stack.push(node2.right);

        } while (!stack.empty());

		return true;
    }
}