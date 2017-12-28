package main

import (
	"fmt"
)

func main() {
	graph := &Graph{[16]*Node{}}
	graph.LoadData("maze.grh")
	graph.PrintGraph()
	fmt.Println("------------------")
	graph.DepthFirstSearch(0, 15, &[]int{})

	// if predNode == -1 {
	// 	fmt.Println("no path")
	// } else {
	// 	fmt.Println(predNode)
	// }
}
