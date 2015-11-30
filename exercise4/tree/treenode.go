package tree

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

// Constructor
func NewTreeNode(val int) *TreeNode {
	return &TreeNode{val, nil, nil}
}
