package main

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	SEP = ","
	NIL = "#"
)

type TreeNode struct {
	Left *TreeNode
	Right *TreeNode
	Val int
}

// 序列化
var treeData string
func Serialize(root *TreeNode) {
	if root == nil {
		treeData += fmt.Sprintf("%s%s", NIL, SEP)
		return
	}
	treeData += fmt.Sprintf("%d%s", root.Val, SEP)
	Serialize(root.Left)
	Serialize(root.Right)
}

//反序列化
var nums []string
func DeSerialize(data string) *TreeNode{
	for _, value := range strings.Split(data, SEP) {
		nums = append(nums, value)
	}
	nums = nums[0:len(nums) - 1]
	return deserialize()
}


func deserialize() *TreeNode{
	if len(nums) == 0 {
		return nil
	}
	first := nums[0]
	fmt.Println(nums)
	nums = nums[1:]
	if first == NIL {
		return nil
	}
	val, _ := strconv.Atoi(first)
	root := &TreeNode{Val: val}
	root.Left = deserialize()
	root.Right = deserialize()
	return root
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
func BuildMaxTree(nums []int, lo, hi int) *TreeNode{
	if lo > hi {
		return nil
	}
	mid := GetMax(nums, lo, hi)
	root := &TreeNode{Val: nums[mid]}
	root.Left = BuildMaxTree(nums, lo, mid - 1)
	root.Right = BuildMaxTree(nums, mid+1, hi)
	return root
}

// 因为append返回的是一个新切片，所以不能用下述通过入参的形式传入，这种情况下外部的slice其实一直是为空的，因为append后是一个新的切片了，与原切片之间相互独立
//func Traverser(root *TreeNode) []int {
//	var slice []int
//	traverser(root, slice)
//	fmt.Println(slice)
//	return slice
//}
//
//func traverser(root *TreeNode, slice []int) {
//	if root == nil {
//		return
//	}
//	slice = append(slice, root.Val)
//	fmt.Println(slice)
//	traverser(root.Left, slice)
//	traverser(root.Right, slice)
//}

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
	root := BuildMaxTree(nums, 0, len(nums) - 1)
	traverser(root)
	fmt.Println(slice)
	slice = slice[:0]
	fmt.Println(slice)
	Serialize(root)
	fmt.Println(fmt.Sprintf("tree : %s", treeData))
	otherRoot := DeSerialize(treeData)
	traverser(otherRoot)
	fmt.Print(slice)
}
