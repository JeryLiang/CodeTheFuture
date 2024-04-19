package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Partition(nums []int, lo, hi int) int {
	p := nums[lo]
	i, j := lo + 1, hi
	for i <= j {
		for i < hi && nums[i] <= p {
			i++
		}
		for j > lo && nums[j] > p {
			j--
		}
		if i >= j {
			break
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	nums[lo], nums[j] = nums[j], nums[lo]
	return j
}

func QuickSort(nums []int, lo, hi int) {
	if lo > hi {
		return
	}
	p := Partition(nums, lo, hi)
	QuickSort(nums, lo, p - 1)
	QuickSort(nums, p + 1, hi)
}

func ShuffleSlice(slice []int) {
	rand.Seed(time.Now().UnixNano()) // 使用当前时间作为随机数种子
	for i := len(slice) - 1; i >0; i-- {
		j := rand.Intn(i + 1) // 生成一个0到i的随机数
		slice[i], slice[j] = slice[j], slice[i] // 交换元素
	}
}

func main(){
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("原始切片:", slice)

	ShuffleSlice(slice)
	fmt.Println("打乱后的切片:", slice)
	QuickSort(slice, 0, len(slice) - 1)
	fmt.Println("排序后的切片: ", slice)
}
