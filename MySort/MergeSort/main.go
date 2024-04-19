package main

import (
	"fmt"
	"math/rand"
	"time"
)

var temp []int

func ShuffleSlice(slice []int) {
	rand.Seed(time.Now().UnixNano()) // 使用当前时间作为随机数种子
	for i := len(slice) - 1; i >0; i-- {
		j := rand.Intn(i + 1) // 生成一个0到i的随机数
		slice[i], slice[j] = slice[j], slice[i] // 交换元素
	}
}

func Merge(nums []int, lo, mid, hi int) {
	copy(temp, nums)
	i, j := lo, mid + 1
	for p := lo; p <= hi; p++ {
		if i == mid + 1 {
			// 左边已排序完
			nums[p] = temp[j]
			j++
		}else if j == hi + 1 {
			nums[p] = temp[i]
			i++
		}else if temp[i] < temp[j] {
			nums[p] = temp[i]
			i++
		}else {
			nums[p] = temp[j]
			j++
		}
	}
}


func MergeSort(nums []int, lo, hi int) {
	// 递归出口lo == hi 说明只有一个元素，不需要再排序了
	if lo == hi {
		return
	}
	mid := lo + (hi - lo) / 2
	MergeSort(nums, lo, mid)
	MergeSort(nums, mid + 1, hi)
	Merge(nums, lo, mid, hi)
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	ShuffleSlice(nums)
	temp = make([]int, len(nums))
	fmt.Println(nums)
	MergeSort(nums, 0, len(nums)-1)
	fmt.Println(nums)
}
