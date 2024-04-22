package main

import "fmt"

type TreeNode struct {
	Left *TreeNode
	Right *TreeNode
	Val int
}

func IsVaildBST(root *TreeNode) bool{
	return isVaildBST(root, nil, nil)
}

// 限定以root为根节点的子树节点必须满足max.Val > root.Val > min.Val
func isVaildBST(root, min, max *TreeNode) bool{
	if root == nil {
		return true
	}
	if min != nil && root.Val <= min.Val {
		return false
	}
	if max != nil && root.Val >= max.Val {
		return false 
	}
	// 限定左子树的最大值是root.Val, 右子树的最小值是root.Val
	return isVaildBST(root.Left, min, root) && isVaildBST(root.Right, root, max)
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
	root.Right.Right = &TreeNode{Val: 7} // true
	// root.Right.Right = &TreeNode{Val: 9} // false
	root.Right.Right.Right = &TreeNode{Val: 8}
	traserver(root)
	fmt.Println(slice)
	fmt.Println(IsVaildBST(root))
}