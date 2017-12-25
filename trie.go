package main

import (
	"fmt"
)

type Node struct {
	Value    string
	Parent   *Node
	Children map[string]*Node
	IsEnded  bool
	Key      string
}

func (node *Node) insert(key string) *Node {
	newNode := &Node{"", node, make(map[string]*Node), false, key}
	node.Children[key] = newNode
	return newNode
}

type Trie struct {
	Root *Node
}

func (trie *Trie) insert(value string) {
	curr := trie.Root
	for _, char := range value {
		curr = curr.insert(string(char))
	}

	curr.IsEnded = true
	curr.Value = value
}

func (trie *Trie) delete(value string) {
	endNode := trie.Root
	for _, char := range value {
		endNode = endNode.Children[string(char)]
		if endNode == nil {
			return
		}
	}

	endNode.IsEnded = false
	trie.deleteNodes(endNode)
}

func (trie *Trie) deleteNodes(node *Node) {
	if !checkIfChildExist(node.Children) && node.Parent != nil {
		parent := node.Parent
		parent.Children[node.Key] = nil
		trie.deleteNodes(parent)
	}
}

func checkIfChildExist(childrens map[string]*Node) bool {
	for _, value := range childrens {
		if value != nil {
			return true
		}
	}

	return false
}

func (trie *Trie) getValue(value string) string {
	curr := trie.Root
	for _, char := range value {
		curr = curr.Children[string(char)]
		if curr == nil {
			return "not found"
		}
	}

	if curr.IsEnded && curr.Value != "" {
		return curr.Value
	}

	return "not found"
}

func main() {
	root := &Node{"", nil, make(map[string]*Node), false, ""}
	trie := &Trie{root}
	testWords := []string{"cat", "category", "dog", "foo", "bar"}
	for _, word := range testWords {
		trie.insert(word)
	}

	trie.delete("cat")
	trie.delete("dog")

	for _, word := range testWords {
		fmt.Println("value for " + word + " is " + trie.getValue(word))
	}
}
