package data_structures

struct type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(value int, left *TreeNode, right *TreeNode) *TreeNode {
	treenode := new(TreeNode)

	treenode.Val = value
	treenode.Left = left
	treenode.Right = right

	return treenode
}