package main

import "fmt"

type Node struct {
	left, right *Node
	value       int
}

func main() {
	originalArr := []int{4, 3, 12, 2, 6, 8, 5, 1}
	originalArr = treeSort(originalArr)
	fmt.Println(originalArr)
}

func (this *Node) addNode(other *Node) {
	if this.value < other.value {
		if this.right == nil {
			this.right = other
		} else {
			this.right.addNode(other)
		}
	}

	if this.value >= other.value {
		if this.left == nil {
			this.left = other
		} else {
			this.left.addNode(other)
		}
	}
}

func createBinarySearchTree(arr []int) *Node {
	root := &Node{nil, nil, arr[0]}
	size := len(arr)
	for i := 1; i < size; i++ {
		root.addNode(&Node{nil, nil, arr[i]})
	}
	return root
}

func treeSort(arr []int) []int {
	root := createBinarySearchTree(arr)
	tempArr := &[]int{}
	insertNode(root, tempArr)
	return *tempArr
}

func insertNode(node *Node, arr *[]int) {
	if node != nil {
		insertNode(node.left, arr)
		*arr = append(*arr, node.value)
		insertNode(node.right, arr)
	}
}
