package main

func main()  {

}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) < 1 || len(postorder) < 1 {
		return nil
	}

	// 1. 找到根节点（后序遍历的最后一个为根结点）
	val := postorder[len(postorder)-1]
	// 2. 找到切割点
	index := findRootIndex(inorder, val)

	return &TreeNode{
		Val: val,
		Left: buildTree(inorder[:index], postorder[:index]),
		Right: buildTree(inorder[index+1:], postorder[index:len(postorder)-1])}
}

func findRootIndex(inorder []int, target int) (index int) {
	for i, x := range inorder {
		if x == target {
			return i
		}
	}
	return -1
}
