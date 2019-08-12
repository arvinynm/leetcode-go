package main

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isMirrorTree(root, root)
}

func isMirrorTree(left *TreeNode, right *TreeNode) bool{
	if left == nil {
		return right == nil
	}
	if right == nil {
		return left == nil
	}
	if left.Val == right.Val {
		return isMirrorTree(left.Left, right.Right) && isMirrorTree(left.Right, right.Left)
	}
	return false
}