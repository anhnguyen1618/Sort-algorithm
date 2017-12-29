package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Node struct {
	IndexPointer int
	NeighborNode *Node
}

type Graph struct {
	Data [16]*Node
}

func (graph *Graph) LoadData(fileName string) {
	byteContent, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Print(err)
	}

	stringContent := string(byteContent)
	lines := strings.Split(stringContent, "\r\n")
	for index, line := range lines {
		numbers := strings.Split(line, " ")

		for i := len(numbers) - 1; i > 0; i-- {
			number, err := strconv.Atoi(numbers[i])
			if err != nil {
				fmt.Println(err)
			}
			graph.Data[index] = &Node{number, graph.Data[index]}
		}
	}
}

func (graph *Graph) PrintGraph() {
	for index, rootNode := range graph.Data {
		stringBuilder := strconv.Itoa(index) + ":"

		currentNode := rootNode
		for currentNode != nil {
			stringBuilder += " " + strconv.Itoa(currentNode.IndexPointer)
			currentNode = currentNode.NeighborNode
		}

		fmt.Println(stringBuilder)
	}
}

func checkIfNodeIsVisited(visited *[]int, node int) bool {
	for _, visitedNode := range *visited {
		if visitedNode == node {
			return true
		}
	}

	return false
}

func (graph *Graph) dfs(from int, to int, visited *[]int, routes *[]int) int {
	if checkIfNodeIsVisited(visited, from) {
		return -1
	}

	if from == to {
		*routes = append(*routes, to)
		return to
	}

	*visited = append(*visited, from)

	currentNode := graph.Data[from]

	for currentNode != nil {
		prevNode := graph.dfs(currentNode.IndexPointer, to, visited, routes)
		if prevNode != -1 {
			*routes = append(*routes, from)
			return currentNode.IndexPointer
		}

		currentNode = currentNode.NeighborNode
	}

	return -1

}

func (graph *Graph) DepthFirstSearch(from int, to int) {
	routes := &[]int{}
	graph.dfs(from, to, &[]int{}, routes)
	fmt.Println("Route from " + strconv.Itoa(from) + " to " + strconv.Itoa(to) + ":")
	stringBuilder := ""
	for i := len(*routes) - 1; i >= 0; i-- {
		if i != 0 {
			stringBuilder += strconv.Itoa((*routes)[i]) + " -> "
		} else {
			stringBuilder += strconv.Itoa((*routes)[i])
		}
	}
	fmt.Println(stringBuilder)
}
