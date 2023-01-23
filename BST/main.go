package BST

import "fmt"

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

var preTree *BST

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

func (tree *BST) Contains(value int) bool {
	currentTree := tree
	for {
		if currentTree.Value == value {
			return true
		}
		if value < currentTree.Value {
			if currentTree.Left == nil {
				break
			}
			currentTree = currentTree.Left
		} else {
			if currentTree.Right == nil {
				break
			}
			currentTree = currentTree.Right
		}
	}
	return false
}

func (tree *BST) FindSmallest() int {
	currentTree := tree
	for {
		if currentTree.Left == nil {
			return currentTree.Value
		}
		currentTree = currentTree.Left
	}
}

func (tree *BST) Remove(value int) *BST {
	currentTree := tree
	if currentTree.Left == nil && currentTree.Right == nil && currentTree.Value == value {
		if preTree != nil && preTree.Right == currentTree {
			preTree.Right = nil
		} else if preTree != nil && preTree.Left == currentTree {
			preTree.Left = nil
		}
		preTree = nil
		return tree
	}
	for {
		if value < currentTree.Value {
			preTree = currentTree
			currentTree = currentTree.Left
		} else if value > currentTree.Value {
			preTree = currentTree
			currentTree = currentTree.Right
		} else {
			if currentTree.Left != nil && currentTree.Right != nil {
				currentTree.Value = currentTree.Right.FindSmallest()
				preTree = currentTree

				currentTree.Right.Remove(currentTree.Value)
			} else if preTree == nil {
				if tree.Right != nil {
					*tree = *tree.Right
					break
				}
				*tree = *tree.Left
				break
			} else if preTree.Left == currentTree {
				if currentTree.Left != nil {
					preTree.Left = currentTree.Left
				} else {
					preTree.Left = currentTree.Right
				}
			} else if preTree.Right == currentTree {
				if currentTree.Left != nil {
					preTree.Right = currentTree.Left
				} else {
					preTree.Right = currentTree.Right
				}
			}

			break
		}
	}
	preTree = nil
	return tree
}

func CreateBst(value int) *BST {
	return &BST{Value: value}
}

func main() {
	root := CreateBst(1)
	root.Insert(2)
	root.Insert(3)
	root.Insert(4)
	root.Remove(1)
	fmt.Println(root)
}
