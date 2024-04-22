package main

import (
	"fmt"
	"math/rand"
	"time"
)

type TreeNode struct {
	Left *TreeNode
	Right *TreeNode
	Val int
}

func InsertIntoBST(root *TreeNode, val int) *TreeNode{
	if root == nil {
		return &TreeNode{Val: val}
	}
	if root.Val > val {
		root.Left = InsertIntoBST(root.Left, val)
	} else {
		root.Right = InsertIntoBST(root.Right, val)
	}
	return root
}

func ShuffleSlice(slice []int) {
	rand.Seed(time.Now().UnixNano()) // 使用当前时间作为随机数种子
	for i := len(slice) - 1; i >0; i-- {
		j := rand.Intn(i + 1) // 生成一个0到i的随机数
		slice[i], slice[j] = slice[j], slice[i] // 交换元素
	}
}


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
	nums := []int{1,2,3,4,5,6}
	ShuffleSlice(nums)
	fmt.Println(nums)
	var root *TreeNode
	root = nil
	for _, num := range nums {
		root = InsertIntoBST(root, num)
	}
	traserver(root)
	fmt.Println(slice)
}