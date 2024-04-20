package main

import "fmt"

type TreeNode struct {
	Left *TreeNode
	Right *TreeNode
	Val int
}

func GetMax(nums []int, lo, hi int) int {
	res := lo
	for i := lo; i <= hi; i++ {
		if nums[i] > nums[res] {
			res = i
		}
	}
	return res
}

// 	构建最大二叉树
func BuilMaxTree(nums []int, lo, hi int) *TreeNode{
	if lo > hi {
		return nil
	}
	mid := GetMax(nums, lo, hi)
	root := &TreeNode{Val: nums[mid]}
	root.Left = BuilMaxTree(nums, lo, mid - 1)
	root.Right = BuilMaxTree(nums, mid+1, hi)
	return root
}



// 前序遍历二叉树
var slice []int
func traverser(root *TreeNode) {
	if root == nil {
		return
	}
	slice = append(slice, root.Val)
	traverser(root.Left)
	traverser(root.Right)
}

func main() {
	nums := []int{3,2,1,6,0,5}
	fmt.Println(nums)
	root := BuilMaxTree(nums, 0, len(nums) - 1)
	traverser(root)
	fmt.Println(slice)
}
