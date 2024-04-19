package main

import "fmt"

type TreeNode struct {
	Left *TreeNode
	Right *TreeNode
	Val int
}

var (
	res int
	depth int
)

// 递归分解问题计算树的最大深度
func TreeMaxDepth(root *TreeNode) int{
	if root == nil {
		return 0
	}
	return max(TreeMaxDepth(root.Left) + 1, TreeMaxDepth(root.Right) + 1)
}

// 遍历二叉树的形式计算树的最大深度
func TreeMaxDepthIter(root *TreeNode) {
	if root == nil {
		// 遍历到了子节点
		res = max(res, depth)
		return
	}
	depth++
	TreeMaxDepthIter(root.Left)
	TreeMaxDepthIter(root.Right)
	depth--
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	root := new(TreeNode)
	root.Right = new(TreeNode)
	root.Left = new(TreeNode)
	root.Right.Right = new(TreeNode)
	n := TreeMaxDepth(root)
	fmt.Println(n)
	TreeMaxDepthIter(root)
	fmt.Println(res)
}
