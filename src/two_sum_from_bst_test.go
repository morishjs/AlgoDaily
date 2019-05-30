package main

import (
	. "go-algo-daily/src/data_structures"
	"testing"
)

func Test_twoSumFromBST(t *testing.T) {
	type args struct {
		root   *TreeNode
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"Target value exists",
			args{root: NewTreeNode(3, NewTreeNode(1, nil, nil), NewTreeNode(8, nil, nil)), target: 11},
			true,
		},
		{
			"Target value does not exists",
			args{root: NewTreeNode(3, NewTreeNode(1, nil, nil), NewTreeNode(8, nil, nil)), target: 2},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSumFromBST(tt.args.root, tt.args.target); got != tt.want {
				t.Errorf("twoSumFromBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
