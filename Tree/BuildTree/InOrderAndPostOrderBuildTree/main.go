package main

import "fmt"

type TreeNode struct {
	Left *TreeNode
	Right *TreeNode
	Val int
}

var ValToIndex map[int]int
// 中序遍历和后序遍历构建二叉树
func InOrderAndPostOrderBuildTree(inorder, postorder []int) *TreeNode{
	ValToIndex = make(map[int]int, len(inorder))
	for index, value := range inorder {
		ValToIndex[value] = index
	}
	return BuildTree(inorder, 0, len(inorder) - 1, postorder, 0, len(postorder) - 1)
}

func BuildTree(inorder []int, lo1, hi1 int, postorder []int, lo2, hi2 int) *TreeNode{
	if lo1 > hi1 || lo2 > hi2 {
		return nil
	}
	value := postorder[hi2]
	index := ValToIndex[value]
	leftSize := index - lo1
	root := &TreeNode{Val: value}
	root.Left = BuildTree(inorder, lo1, lo1 + leftSize - 1, postorder, lo2, lo2 + leftSize - 1)
	root.Right = BuildTree(inorder, lo1 + leftSize + 1, hi1, postorder, lo2 + leftSize, hi2 - 1)
	return root
}


var (
	slice1 []int // 存中序遍历结果
	slice2 []int // 存后序遍历结果
)
func traserver(root *TreeNode) {
	if root == nil {
		return
	}
	traserver(root.Left)
	slice1 = append(slice1, root.Val)
	traserver(root.Right)
	slice2 = append(slice2, root.Val)
}

func main() {
	inorder := []int{9,3,15,20,7}
	postorder := []int{9,15,7,20,3}
	root := InOrderAndPostOrderBuildTree(inorder, postorder)
	traserver(root)
	fmt.Println(slice1)
	fmt.Println(slice2)
}