//671. Second Minimum Node In a Binary Tree
/*

Given a non-empty special binary tree consisting of nodes with the non-negative value,
where each node in this tree has exactly two or zero sub-node. If the node has two sub-nodes,
then this node's value is the smaller value among its two sub-nodes. More formally,
the property root.val = min(root.left.val, root.right.val) always holds.

Given such a binary tree, you need to output the second minimum value in the set made of all the nodes' value in the whole tree.

If no such second minimum value exists, output -1 instead.

*/

package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var min1 int
var ans int = math.MaxInt32

func findSecondMinimumValue(root *TreeNode) int {
	min1 = root.Val
	dfs(root)

	if ans < math.MaxInt32 {
		return ans
	}

	return -1
}

func dfs(root *TreeNode) {
	if root == nil {
		return
	}

	//this is because root.val = min (left.val, right.val)
	if min1 < root.Val && root.Val < ans {
		ans = root.Val
	} else if min1 == root.Val {
		dfs(root.Left)
		dfs(root.Right)
	}

}

func main() {

	n1 := TreeNode{3, nil, nil}
	n2 := TreeNode{2, nil, nil}
	root := TreeNode{2, &n1, &n2}

	fmt.Println(findSecondMinimumValue(&root))

}
