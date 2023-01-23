package BSTValidate

import (
	"fmt"
	"math"
)

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

func (tree *BST) Insert(value int) *BST {
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

func (tree *BST) ValidateNode(max int, min int) bool {
	if tree == nil {
		return true
	}
	if tree.Value < min || tree.Value >= max {
		return false
	}
	leftIsValid := tree.Left.ValidateNode(tree.Value, min)
	return leftIsValid && tree.Right.ValidateNode(max, tree.Value)
}

func (tree *BST) ValidateBst() bool {
	return tree.ValidateNode(math.MaxInt, math.MinInt)
}

func CreateBst(value int) *BST {
	return &BST{Value: value}
}

func main() {
	root := CreateBst(10)
	root.Insert(5)
	root.Insert(2)
	root.Insert(1)
	root.Insert(5)
	root.Insert(15)
	root.Insert(22)

	fmt.Println(root.ValidateNode(math.MaxInt, math.MinInt))
}
