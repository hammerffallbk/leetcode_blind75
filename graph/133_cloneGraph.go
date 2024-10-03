package graph

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	visited := make(map[int]*Node)
	var dfs func(node *Node) *Node
	dfs = func(node *Node) *Node {
		if node == nil {
			return nil
		}
		if n, ok := visited[node.Val]; ok {
			return n
		}
		clone := &Node{Val: node.Val, Neighbors: make([]*Node, 0, len(node.Neighbors))}
		visited[node.Val] = clone
		for _, n := range node.Neighbors {
			clone.Neighbors = append(clone.Neighbors, dfs(n))
		}
		return clone
	}
	return dfs(node)
}
