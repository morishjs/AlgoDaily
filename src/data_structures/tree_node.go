package ds

type TreeNode struct {
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

func (treenode *TreeNode) Find(val int) bool {
	if treenode.Val == val {
		return true
	}

	if treenode.Left != nil {
		found := treenode.Left.Find(val)
		if found {
			return true
		}
	}

	if treenode.Right != nil {
		found := treenode.Right.Find(val)
		if found {
			return true
		}
	}
	
	return false
}