package main

import "fmt"

type TreeNode struct {
	Left *TreeNode
	Right *TreeNode
	Val int
}

var (
	res int
	resNode *TreeNode
)

func TheKNum(root *TreeNode, k int){
	traserver(root, k)
}

// 中序遍历二叉搜索树得到的结果为升序，如果想要得到降序，则需要先遍历右子树再到左子树
func traserver(root *TreeNode, k int) {
	if root == nil {
		return
	}
	traserver(root.Left, k)
	res++
	if res == k {
		resNode = root
		return 
	}
	traserver(root.Right, k)
	
}

func main() {
	root := &TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 3}
	root.Right = &TreeNode{Val: 6}
	root.Left.Left = &TreeNode{Val: 2}
	root.Left.Right = &TreeNode{Val: 4}
	root.Left.Left.Left = &TreeNode{Val: 1}
	TheKNum(root, 3)
	fmt.Println(resNode.Val)
}