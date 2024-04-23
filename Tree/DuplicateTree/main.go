package main

import "fmt"

type TreeNode struct {
	Left *TreeNode
	Right *TreeNode
	Val int
}

// 输出一颗树的重复子树对应的根节点（对应力扣652题：寻找重复的子树）
var (
	memo map[string]int
	duplicateList []*TreeNode
)
func FindDuplicateSubTree(root *TreeNode) {
	memo = make(map[string]int, 10)
	traverser(root)
}

var slice []int
func traverser(root *TreeNode) string{
	if root == nil {
		return "#"
	}
	slice = append(slice, root.Val) // 前序遍历结果集，为了方便观察树的情况
	left := traverser(root.Left)
	right := traverser(root.Right)
	subTree := fmt.Sprintf("%s,%s,%d", left, right, root.Val)
	value, _ := memo[subTree] // 如果取不到map会返回value类型的默认零值，此处为int,为此取不到的情况下会返回0
	memo[subTree] = value + 1
	if value == 1 {
		duplicateList = append(duplicateList, root)
	}
	return subTree
}

func main() {
	root := &TreeNode{Val: 1, Right: &TreeNode{Val: 3}, Left: &TreeNode{Val: 2}}
	root.Left.Left = &TreeNode{Val: 4}
	root.Right.Left = &TreeNode{Val: 2}
	root.Right.Right = &TreeNode{Val: 4}
	root.Right.Left.Left = &TreeNode{Val: 4}
	FindDuplicateSubTree(root)
	fmt.Println(slice)
	for _, node := range duplicateList {
		fmt.Println(node.Val)
	}
}