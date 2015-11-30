package main

import (
	"./tree"
	"fmt"
)

func main() {
	// tree1 is balanced
	//       5
	//      / \
	//     2   7
	//    /
	//   1
	//
	tree1 := tree.NewTreeNode(5)
	tree1.Left = tree.NewTreeNode(2)
	tree1.Right = tree.NewTreeNode(7)
	tree1Left := tree1.Left
	tree1Left.Left = tree.NewTreeNode(1)

	fmt.Printf("Tree1 is balanced or not: %v\n", tree.IsBalanced(tree1))

	// tree2 is not balanced
	//       5
	//      /
	//     2
	//    /
	//   1
	//
	tree2 := tree.NewTreeNode(5)
	tree2.Left = tree.NewTreeNode(2)
	tree2Left := tree2.Left
	tree2Left.Left = tree.NewTreeNode(1)

	fmt.Printf("Tree2 is balanced or not: %v\n", tree.IsBalanced(tree2))
}
