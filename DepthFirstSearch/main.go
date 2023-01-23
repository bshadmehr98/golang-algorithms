package DepthFirstSearch

import (
	"fmt"
)

type Node struct {
	Name     string
	Children []*Node
}

func (n *Node) DepthFirstSearch(array []string) []string {
	array = append(array, n.Name)
	for _, node := range n.Children {
		array = node.DepthFirstSearch(array)
	}
	return array
}

func (n *Node) Insert(value string) *Node {
	node := CreateNode(value)
	n.Children = append(n.Children, node)
	return node
}

func CreateNode(value string) *Node {
	return &Node{Name: value}
}

func main() {
	a := CreateNode("A")
	b := a.Insert("B")
	b.Insert("E")
	f := b.Insert("F")
	f.Insert("I")
	f.Insert("J")
	a.Insert("C")
	d := a.Insert("D")
	d.Insert("H")
	g := d.Insert("G")
	g.Insert("K")

	arr := make([]string, 0)
	fmt.Println(a.DepthFirstSearch(arr))
}
