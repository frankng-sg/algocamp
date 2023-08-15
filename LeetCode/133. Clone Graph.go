package main

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	var clone func(*Node) *Node
	visited := make(map[*Node]*Node)
	clone = func(node *Node) *Node {
		var newNode Node
		visited[node] = &newNode
		newNode.Val = node.Val
		for _, v := range node.Neighbors {
			if visited[v] == nil {
				newNode.Neighbors = append(newNode.Neighbors, clone(v))
			} else {
				newNode.Neighbors = append(newNode.Neighbors, visited[v])
			}
		}
		return &newNode
	}
	if node == nil {
		return nil
	}
	return clone(node)
}

func main() {
	node1 := Node{Val: 1}
	node2 := Node{Val: 2}
	node3 := Node{Val: 3}
	node1.Neighbors = []*Node{&node2, &node3}
	node2.Neighbors = []*Node{&node1, &node3}
	node3.Neighbors = []*Node{&node1, &node2}
	cloneGraph(&node1)
}
