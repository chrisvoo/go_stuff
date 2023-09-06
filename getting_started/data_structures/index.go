package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

func Print(root *Node) {
	nodeWalk := root
	for nodeWalk.Next != nil {
		fmt.Println(nodeWalk.Value)
		nodeWalk = nodeWalk.Next
	}
	fmt.Println(nodeWalk.Value)
}

func main() {
	aa := Node{Value: 1}
	bb := Node{Value: 2}
	cc := Node{Value: 3}
	aa.Next = &bb
	bb.Next = &cc

	fmt.Println(aa.Next) // &{2 0xc000014080}
	Print(aa.Next)
}
