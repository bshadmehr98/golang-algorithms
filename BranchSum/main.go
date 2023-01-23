package main

import (
	"fmt"
)

type BinaryTree struct {
	Value int
	Left  *BinaryTree
	Right *BinaryTree
}

func (tree *BinaryTree) Insert(value int) *BinaryTree {
	currentTree := tree
	for {
		if value < currentTree.Value {
			if currentTree.Left == nil {
				currentTree.Left = CreateBst(value)
				break
			}
			currentTree = currentTree.Left
		} else {
			if currentTree.Right == nil {
				currentTree.Right = CreateBst(value)
				break
			}
			currentTree = currentTree.Right
		}
	}
	return tree
}

func (tree *BinaryTree) InOrderTraverse(array []int, s int) []int {
	if tree.Right == nil && tree.Left == nil {
		array = append(array, s+tree.Value)
		return array
	}
	if tree.Left != nil {
		array = tree.Left.InOrderTraverse(array, s+tree.Value)
	}
	//array = append(array, tree.Value)
	if tree.Right != nil {
		array = tree.Right.InOrderTraverse(array, s+tree.Value)
	}
	return array
}

func BranchSums(root *BinaryTree) []int {
	arr := make([]int, 0)
	return root.InOrderTraverse(arr, 0)
}

func CreateBst(value int) *BinaryTree {
	return &BinaryTree{Value: value}
}

func main() {
	root := CreateBst(10)
	root.Insert(5)
	root.Insert(2)
	root.Insert(1)
	root.Insert(5)
	root.Insert(15)
	root.Insert(22)

	fmt.Println(BranchSums(root))
}
