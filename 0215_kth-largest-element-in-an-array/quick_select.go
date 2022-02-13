package main

import "math"

func findKthLargest2(nums []int, k int) int {
	return qselect(nums, 0, len(nums)-1, k)
}
func qselect(nums []int, left, right, k int) int {
	for {
		if left == right {
			return nums[left]
		}
		pivotIdx := int(math.Floor(float64(left+right) / 2))
		pivotIdx = partition(nums, left, right, pivotIdx)
		targetPos := len(nums) - k
		if targetPos == pivotIdx {
			return nums[targetPos]
		} else if targetPos < pivotIdx {
			right = pivotIdx - 1
		} else {
			left = pivotIdx + 1
		}
	}
}

func partition(nums []int, left, right, pivotIdx int) int {
	pivot := nums[pivotIdx]
	nums[pivotIdx], nums[right] = nums[right], nums[pivotIdx]

	storeIdx := left
	for i := left; i < right; i++ {
		if nums[i] < pivot {
			nums[storeIdx], nums[i] = nums[i], nums[storeIdx]
			storeIdx++
		}
	}
	nums[storeIdx], nums[right] = nums[right], nums[storeIdx]
	return storeIdx
}
