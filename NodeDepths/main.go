package NodeDepths

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

func NodeDepthHelper(root *BinaryTree, sum int, curDis int) int {
	if root.Left == nil && root.Right == nil {
		return sum + curDis
	}
	if root.Left != nil {
		sum = NodeDepthHelper(root.Left, sum, curDis+1)
	}
	sum += curDis
	if root.Right != nil {
		sum = NodeDepthHelper(root.Right, sum, curDis+1)
	}
	return sum
}

func NodeDepths(root *BinaryTree) int {
	return NodeDepthHelper(root, 0, 0)
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

	fmt.Println(NodeDepths(root))
}
