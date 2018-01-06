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
	Weight       int
}

type Graph struct {
	Data [16]*Node
}

type Row struct {
	DistanceFromStart int
	PreviousNode      int
}

func (row Row) changeDistanceFromStart(newDistance int) {
	row.DistanceFromStart = newDistance
}

func (row Row) changePreviousNodeIndex(newNodeIndex int) {
	row.PreviousNode = newNodeIndex
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

			oldRootNode := graph.Data[index]
			graph.Data[index] = &Node{number, oldRootNode, 1}
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

func filterOut(unvisited *[]int, node int) {
	filteredOutput := []int{}
	for _, value := range *unvisited {
		if value != node {
			filteredOutput = append(filteredOutput, value)
		}
	}
	*unvisited = filteredOutput
}

func (graph *Graph) PreFillTrackTable(from int) map[int]Row {
	trackTable := make(map[int]Row)

	for index := range graph.Data {
		if index == from {
			trackTable[index] = Row{0, -1}
		} else {
			trackTable[index] = Row{999999999999999999, -1}
		}
	}
	return trackTable
}

func (graph *Graph) Djikstra(from int, to int) {
	unVisited := []int{from}
	visited := []int{}

	trackTable := graph.PreFillTrackTable(from)

	currentIndex := from

	for len(unVisited) != 0 {
		filterOut(&unVisited, currentIndex)

		visited = append(visited, currentIndex)
		currentNode := graph.Data[currentIndex]
		distanceMappings := make(map[int]int)

		distanceFromStartToRootNode := trackTable[currentIndex].DistanceFromStart

		for currentNode != nil {
			distanceFromStart := distanceFromStartToRootNode
			if currentNode.IndexPointer != from {
				distanceFromStart = distanceFromStartToRootNode + currentNode.Weight
			}

			if !checkIfNodeIsVisited(&visited, currentNode.IndexPointer) {
				distanceMappings[currentNode.IndexPointer] = distanceFromStart
				unVisited = append(unVisited, currentNode.IndexPointer)
			}

			if distanceFromStart < trackTable[currentNode.IndexPointer].DistanceFromStart {
				temp := trackTable[currentNode.IndexPointer]
				temp.DistanceFromStart = distanceFromStart
				trackTable[currentNode.IndexPointer] = temp

				temp2 := trackTable[currentNode.IndexPointer]
				temp2.PreviousNode = currentIndex
				trackTable[currentNode.IndexPointer] = temp2
			}
			currentNode = currentNode.NeighborNode
		}

		smallestNode := currentIndex
		smallestDistance := 999999999999999999

		for key, value := range distanceMappings {
			if value < smallestDistance {
				smallestNode, smallestDistance = key, value
			}
		}

		if currentIndex == smallestNode {
			if len(unVisited) == 0 {
				break
			}
			currentIndex = unVisited[len(unVisited)-1]
		} else {
			currentIndex = smallestNode
		}

	}

	fmt.Println(trackTable)
}
