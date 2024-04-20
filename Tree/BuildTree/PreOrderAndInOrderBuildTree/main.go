package main

import "fmt"

type TreeNode struct {
	Left *TreeNode
	Right *TreeNode
	Val int
}

// 前序遍历和中序遍历构建二叉树
var inPos map[int]int
func PreOrderAndInOrderBuildTree(preorder []int, inorder []int) *TreeNode{
	inPos = make(map[int]int, len(inorder))
	for index, value := range inorder {
		inPos[value] = index
	}
	return BuildTree(preorder, 0, len(preorder) - 1, inorder, 0, len(inorder) - 1)
}

func BuildTree(preorder []int, lo1, hi1 int, inorder []int, lo2, hi2 int) *TreeNode{
	if lo1 > hi1 || lo2 > hi2 {
		return nil
	}
	value := preorder[lo1]
	index := inPos[value]
	root := &TreeNode{Val: value}
	// 其中index - lo2 为左子树元素个数
	root.Left = BuildTree(preorder, lo1 + 1, lo1 + index- lo2, inorder, lo2, index - 1)
	root.Right = BuildTree(preorder, lo1 + 1 + index - lo2, hi1, inorder, index + 1, hi2)
	return root
}

var slice1 []int
var slice2 []int
func traverser(root *TreeNode) {
	if root == nil {
		return
	}
	// 前序
	slice1 = append(slice1, root.Val)
	traverser(root.Left)
	// 中序
	slice2 = append(slice2, root.Val)
	traverser(root.Right)
}

func main() {
	preorder := []int{3,9,20,15,7}
	inorder := []int{9,3,15,20,7}
	root := PreOrderAndInOrderBuildTree(preorder, inorder)
	traverser(root)
	fmt.Println("前序遍历结果：", slice1)
	fmt.Println("中序遍历结果", slice2)
}