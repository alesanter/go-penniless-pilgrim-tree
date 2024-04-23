package main

import "fmt"

type Node struct {
	Row int
	Col int
}

func (n *Node) String() string {
	return fmt.Sprintf("{%d %d}", n.Row, n.Col)
}

var AllNodes map[int]map[int]*Node

func InitNode(s int) {
	AllNodes = make(map[int]map[int]*Node, s)
	for row := 0; row < s; row++ {
		AllNodes[row] = make(map[int]*Node, s)
		for col := 0; col < s; col++ {
			AllNodes[row][col] = &Node{row, col}
		}
	}
}

// Nullable
func GetNode(r, c int) *Node {
	ns, ok := AllNodes[r]
	if !ok {
		return nil
	}
	n, ok := ns[c]
	if !ok {
		return nil
	}
	return n
}

// Nullable <- GetNode
func (n *Node) Move(d Dir) *Node {
	switch d {
	case DIR_RIGHT:
		return GetNode(n.Row, n.Col+1)
	case DIR_LEFT:
		return GetNode(n.Row, n.Col-1)
	case DIR_UP:
		return GetNode(n.Row-1, n.Col)
	case DIR_DOWN:
		return GetNode(n.Row+1, n.Col)
	default:
		panic(d)
	}
}

func (n *Node) Within(s int) bool {
	return 0 <= n.Row && n.Row < s && 0 <= n.Col && n.Col < s
}
