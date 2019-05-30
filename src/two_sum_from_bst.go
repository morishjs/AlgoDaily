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

// Solution: https://algodaily.com/challenges/two-sum-from-bst

package main

import (
	. "go-algo-daily/src/data_structures"
)

func twoSumFromBST(root *TreeNode, target int) bool {
	sortedArr := root.InOrder()

	left, right := 0, len(sortedArr)-1

	for left < right {
		twoSum := sortedArr[left] + sortedArr[right]

		if twoSum == target {
			return true
		}

		if twoSum < target {
			left += 1
		} else {
			right -= 1
		}
	}

	return false
}
