package main

import "fmt"

type Node struct {
	data  int
	left  *Node
	right *Node
}

func main() {
	nodes := []int{1, 2, 4, -1, -1, 5, -1, -1, 3, -1, 6, -1, -1}

	idx := -1
	root := buildBinaryTreeFromPreOrder(nodes, &idx)
	printBinaryTree(root)
	fmt.Println()
}

func printBinaryTree(root *Node) {
	if root == nil {
		return
	}

	fmt.Print(root.data, " ")
	printBinaryTree(root.left)
	printBinaryTree(root.right)
}

func buildBinaryTreeFromPreOrder(nodes []int, idx *int) *Node {
	*idx++

	if nodes[*idx] == -1 {
		return nil
	}

	newNode := &Node{data: nodes[*idx]}
	newNode.left = buildBinaryTreeFromPreOrder(nodes, idx)
	newNode.right = buildBinaryTreeFromPreOrder(nodes, idx)
	return newNode
}
