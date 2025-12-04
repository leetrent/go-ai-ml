package main

import (
	"errors"
	"fmt"
)

type DepthFirstSearch struct {
	Frontier []*Node
	Game     *Maze
}

func (dfs *DepthFirstSearch) GetFrontier() []*Node {
	return dfs.Frontier
}

func (dfs *DepthFirstSearch) add(i *Node) {
	// last-in / first-out
	dfs.Frontier = append(dfs.Frontier, i)
}

func (dfs *DepthFirstSearch) ContainsState(i *Node) bool {
	for _, x := range dfs.Frontier {
		if x.State == i.State {
			return true
		}
	}
	return false
}

func (dfs *DepthFirstSearch) Empty() bool {
	return len(dfs.Frontier) == 0
}

func (dfs *DepthFirstSearch) Remove() (*Node, error) {
	if len(dfs.Frontier) > 0 {
		if dfs.Game.Debug {
			fmt.Println("Frontier before remove:")
			for _, x := range dfs.Frontier {
				fmt.Println("Node:", x.State)
			}
		}
		node := dfs.Frontier[len(dfs.Frontier)-1]
		// deletes last element from slice
		dfs.Frontier = dfs.Frontier[:len(dfs.Frontier)-1]
		return node, nil
	}

	return nil, errors.New("frontier is empty")
}
