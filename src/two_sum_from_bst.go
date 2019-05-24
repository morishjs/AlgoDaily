// This problem is a fun combination of the famous Two Sum problem combined with with a binary search tree. Given the root of a binary search tree, and a target integer K, return two nodes in said tree whose sum equals K. If no pair fits this match, return false.

// Say we are given this tree:

// /***

//   Input:
//     5    
//    / \     
//   3   8
//  /        
// 1

// ***/
// Each node is defined in the following manner:

// class Node {
//     constructor(val) {
//         this.value = val;
//         this.left = null;
//         this.right = null;
//     }
// }
// And thus the below setup and execution code should hold true:

// const root = new Node(5);
// root.left = new Node(3);
// root.right = new Node(8);
// root.left.left = new Node(1);

// const target = 11;

// twoSumFromBST(root, target);
// // true because 3 + 8 = 11

package main

import ds "go-algo-daily/src/data_structures";

func main()  {
	root = ds.NewTreeNode(5, ds.NewTreeNode(3, ds.NewTreeNode(1)), ds.NewTreeNode(8))

	target = 11

	twoSumFromBST(root, target)
}

func twoSumFromBST(root *ds.TreeNode, target int) {
	if target - root.Val == 0 {
		fmt.PrintF("We found it!")
		return
	}

	rest := target - root.Val

	if rest > root.Val {
		twoSumFromBST(root.Right, rest)
	} else {
		twoSumFromBST(root.Left, rest)
	}
}