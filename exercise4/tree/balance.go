package tree

func IsBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	// Check height difference of left child tree and right child tree
	// Return false is difference is more than 1
	diff := getHeight(root.Left) - getHeight(root.Right)
	if diff > 1 || diff < -1 {
		return false
	}

	// Check IsBalanced for child trees/
	return IsBalanced(root.Left) && IsBalanced(root.Right)
}

// Get height of the tree
func getHeight(root *TreeNode) (height int) {
	if root == nil {
		height = 0
		return
	}

	leftHeight := getHeight(root.Left)
	rightHeight := getHeight(root.Right)
	// Height of current treenode is height of tallest child plus 1.
	if leftHeight > rightHeight {
		height = leftHeight + 1
	} else {
		height = rightHeight + 1
	}

	return
}
