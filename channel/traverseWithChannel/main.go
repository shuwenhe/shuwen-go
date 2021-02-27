package main

type Node struct {
	val         int
	left, right *Node
}

func main() {
	root := &Node{
		val: 3,
	}
	c := root.TraverseWithChannel()
	maxNode := 0
	for _, node := range c {
		if node.Value > maxNode {
			maxNode = node.Value
		}
	}
}

func (node *Node) TraverseWithChannel() chan *Node {
}
