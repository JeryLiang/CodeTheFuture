package main

import "fmt"

type TreeNode struct {
	Left *TreeNode
	Right *TreeNode
	Val int
}

// 力扣538题：二叉搜索树转换为累加树
// 利用中序遍历二叉搜索树即为升序，将先遍历Left改为先遍历Right则为降序
var sum int
var slice []int 
var slice2 []int
func traserver(root *TreeNode) {
	if root == nil {
		return
	}
	traserver(root.Right)
	slice = append(slice, root.Val) // 与该题本身无关,实际做题时应去掉
	sum += root.Val
	root.Val = sum
	slice2 = append(slice2, root.Val) //与该题本身无关，实际做题时应去掉
	traserver(root.Left)
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
	fmt.Println(slice2)
}