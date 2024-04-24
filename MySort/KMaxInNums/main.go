package main

import (
	"fmt"
	"time"
	"math/rand"
)

// 力扣215题[数组中的第K个最大元素] 方法一：利用快速选择排序算法
// 题目问「第 k 个最大的元素」，相当于数组升序排序后「排名第 n - k 的元素」，为了方便表述，后文另 k' = n - k。
// 如何知道「排名第 k' 的元素」呢？其实在快速排序算法 partition 函数执行的过程中就可以略见一二。
// partition 函数会将 nums[p] 排到正确的位置，使得 nums[lo..p-1] < nums[p] < nums[p+1..hi]：
// 这时候，虽然还没有把整个数组排好序，但我们已经让 nums[p] 左边的元素都比 nums[p] 小了，也就知道 nums[p] 的排名了
func FindKthLargest(nums []int, k int) int {
	ShuffleSlice(nums)
	lo, hi := 0, len(nums) - 1
	// 转化排名为n-k的元素
	k = len(nums) - k
	for lo <= hi {
		p := Partition(nums, lo, hi)
		if p < k {
			lo = p + 1
		} else if p > k {
			hi = p - 1
		} else {
			// 找到第n-k元素
			return nums[p]
		}
	}
	return -1
}


func Partition(nums []int, lo, hi int) int{
	i, j := lo + 1, hi
	num := nums[lo]
	for i <= j {
		for i <= hi && nums[i] < num {
			i++
		}
		for j >= lo + 1 && nums[j] > num {
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

// 快速排序
func QuickSort(nums []int, lo, hi int) {
	if lo >= hi {
		return
	}
	p := Partition(nums, lo, hi)
	QuickSort(nums, lo, p)
	QuickSort(nums, p+1, hi)
}

func ShuffleSlice(nums []int) {
	rand.Seed(time.Now().UnixNano())
	for i := len(nums) - 1; i >= 0; i-- {
		n := rand.Intn(i + 1)
		nums[i], nums[n] = nums[n], nums[i]
	}
}


// 力扣215题[数组中的第K个最大元素] 方法二：利用小顶堆
func FindKthLargest2(nums []int, k int) int {
    heapSize := len(nums)
    buildMaxHeap(nums, heapSize)
    for i := len(nums) - 1; i >= len(nums) - k + 1; i-- {
        nums[0], nums[i] = nums[i], nums[0]
        heapSize--
        maxHeapify(nums, 0, heapSize)
    }
    return nums[0]
}

func buildMaxHeap(a []int, heapSize int) {
    for i := heapSize/2; i >= 0; i-- {
        maxHeapify(a, i, heapSize)
    }
}

func maxHeapify(a []int, i, heapSize int) {
    l, r, largest := i * 2 + 1, i * 2 + 2, i
    if l < heapSize && a[l] > a[largest] {
        largest = l
    }
    if r < heapSize && a[r] > a[largest] {
        largest = r
    }
    if largest != i {
        a[i], a[largest] = a[largest], a[i]
        maxHeapify(a, largest, heapSize)
    }
}

func main() {
	nums := []int{1,2,3,4,5,6,7,8,9}
	ShuffleSlice(nums)
	fmt.Println(nums)
	QuickSort(nums, 0, len(nums)-1)
	fmt.Println(nums)
	fmt.Println(FindKthLargest2(nums, 5))
}