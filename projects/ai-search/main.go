package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

// Set ids for search types.
const (
	DFS = iota
	BFS
	GBFS
	ASTAR
	DIJKSTRA
)

// main is the entry point to our application.
func main() {
	// Declare some variables.
	var m Maze
	var maze, searchType string
	var debugMode bool

	// Read command line flags, and set some sensible defaults.
	flag.StringVar(&maze, "file", "maze.txt", "maze file")
	flag.StringVar(&searchType, "search", "dfs", "search type")
	flag.BoolVar(&debugMode, "debug", false, "debug mode")
	flag.Parse()

	// Load and parse the maze file.
	err := m.Load(maze, debugMode)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	startTime := time.Now()

	switch searchType {
	case "dfs":
		m.SearchType = DFS
		solveDFS(&m)
	default:
		fmt.Println("Invalid search type")
		os.Exit(1)
	}

	if len(m.Solution.Actions) > 0 {
		fmt.Println("Solution:")
		m.printMaze()
		fmt.Println("Solution is", len(m.Solution.Cells), "steps.")
		fmt.Println("Time to solve:", time.Since(startTime))
	} else {
		fmt.Println("No solution.")
	}

	fmt.Println("Explored", len(m.Explored), "nodes.")
}
