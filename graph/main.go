package main

import (
	"fmt"
)

func main() {
	graph := &Graph{[16]*Node{}}
	graph.LoadData("maze.grh")
	graph.PrintGraph()
	fmt.Println("------------------")
	// graph.DepthFirstSearch(0, 15)
	graph.Djikstra(12, 5)
}
