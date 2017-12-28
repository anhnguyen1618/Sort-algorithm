package main

func main() {
	graph := &Graph{[16]*Node{}}
	graph.LoadData("maze.grh")
	graph.PrintGraph()
}
