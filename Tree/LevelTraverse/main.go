package main

import "fmt"

type TreeNode struct {
	Left *TreeNode
	Right *TreeNode
	Val int
}

var LevelNums []int

func LevelTraverse(root *TreeNode) {
	if root == nil {
		return
	}
	var slice []*TreeNode
	slice = append(slice, root)
	for len(slice) > 0 {
		num := len(slice)
		for i := 0; i < num; i++ {
			cur := slice[0]
			LevelNums = append(LevelNums, cur.Val)
			if cur.Left != nil {
				slice = append(slice, cur.Left)
			}
			if cur.Right != nil {
				slice = append(slice, cur.Right)
			}
			slice = slice[1:]
		}
	}
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}},
		Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 5}},
	}
	LevelTraverse(root)
	fmt.Println(LevelNums)
}
