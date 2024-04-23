package main

import "fmt"

type TreeNode struct {
	Left *TreeNode
	Right *TreeNode
	Val int
}

func GetMin(root *TreeNode) *TreeNode{
	for root.Left != nil {
		root = root.Left
	}
	return root
}

// 删除BST二叉搜索树中的节点
func DeleteNodeBST(root *TreeNode, key int) *TreeNode{
	if root.Val == key {
		// 找到key了需要删除
		if root.Left == nil && root.Right == nil {
			// 左右子树都为空，直接删除（返回nil, 不再返回root）
			return nil
		} else if root.Left == nil {
			// 右子树不为空，直接替换要删除的节点
			return root.Right
		} else if root.Right == nil {
			// 左子树不为空，直接替换要删除的节点
			return root.Left
		} else {
			// 左右子树都非空，需要在左子树中找到最大的节点或者右子树中最小的节点来接替被删除节点
			// 此处以右子树最小节点来接替删除节点
			minNode := GetMin(root.Right)
			// 修改值（一般比较复杂的数据结构是没办法直接替换值来实现交换节点的，此处只是展示删除节点）
			root.Val = minNode.Val
			// 转为去删除minNode节点
			root.Right = DeleteNodeBST(root.Right, minNode.Val)
		}
	} else if root.Val > key {
		// 去左子树找
		root.Left = DeleteNodeBST(root.Left, key)
	} else if root.Val < key {
		// 去右子树找
		root.Right = DeleteNodeBST(root.Right, key)
	}
	return root
}

// 插入数据到二叉搜索树
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

// 中序遍历
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
	nums := []int{5,2,6,1,4,7,3}
	root := &TreeNode{}
	root = nil
	for _, num := range nums {
		root = InsertIntoBST(root, num)
	}
	traserver(root)
	fmt.Println(slice)
	// 清除slice存储的数据
	slice = slice[:0]
	// 删除节点6
	DeleteNodeBST(root, 6)
	traserver(root)
	fmt.Println(slice)
	slice = slice[:0]
	// 删除节点2
	DeleteNodeBST(root, 2)
	traserver(root)
	fmt.Println(slice)
	slice = slice[:0]
}