package ds

import (
	"testing"
)

func TestTreeNode_Find(t *testing.T) {
	type fields struct {
		Val   int
		Left  *TreeNode
		Right *TreeNode
	}
	type args struct {
		val int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			"Find a node",
			fields{5, NewTreeNode(3, NewTreeNode(1, nil, nil), nil), NewTreeNode(8, nil, nil)},
			args{8},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			treenode := &TreeNode{
				Val:   tt.fields.Val,
				Left:  tt.fields.Left,
				Right: tt.fields.Right,
			}
			if got := treenode.Find(tt.args.val); got != tt.want {
				t.Errorf("TreeNode.Find() = %v, want %v", got, tt.want)
			}
		})
	}
}
