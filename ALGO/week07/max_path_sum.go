package main

import "math"

func main() {

}

//TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt32

	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := dfs(root.Left)
		right := dfs(root.Right)

		innerMaxSum := left + root.Val + right
		maxSum = max(maxSum, innerMaxSum)
		outputMaxSum := root.Val + max(left, right) // left,right都是非负的，就不用和0比较了
		return max(outputMaxSum, 0)
	}

	dfs(root)
	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
