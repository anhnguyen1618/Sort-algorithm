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
			// fmt.Println(numbers[i])
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
