package main

import "fmt"

type TreeNode struct {
	Left *TreeNode
	Right *TreeNode
	Val int
}

func FindNumInBST(root *TreeNode, target int) bool{
	if root == nil {
		return false
	}
	if root.Val == target {
		return true
	}else if root.Val < target {
		return FindNumInBST(root.Right, target)
	}else {
		return FindNumInBST(root.Left, target)
	}

}

// 二叉搜索树中序遍历为升序
var slice []int
func traserver(root *TreeNode) {
	if root == nil {
		return
	}
	traserver(root.Left)
	slice = append(slice, root.Val)
	traserver(root.Right)
}

func main() {
	root := &TreeNode{Val: 4}
	root.Left = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 6}
	root.Left.Left = &TreeNode{Val: 0}
	root.Left.Right = &TreeNode{Val: 2}
	root.Left.Right.Right = &TreeNode{Val: 3}
	root.Right.Left = &TreeNode{Val: 5}
	root.Right.Right = &TreeNode{Val: 7}
	root.Right.Right.Right = &TreeNode{Val: 8}
	traserver(root)
	fmt.Println(slice)
	fmt.Println(FindNumInBST(root, 3))
}