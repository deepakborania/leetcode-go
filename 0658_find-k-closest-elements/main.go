package main

func findClosestElements(arr []int, k int, x int) []int {
	if k == len(arr) {
		return arr
	}
	center := findClosestElementIdx(arr, x)
	left := center
	right := center
	for right-left+1 != k {
		if left-1 >= 0 && right+1 < len(arr) {
			if abs(arr[left-1]-x) <= abs(arr[right+1]-x) {
				left--
			} else {
				right++
			}
		} else if right+1 >= len(arr) {
			left--
		} else {
			right++
		}
	}
	return arr[left : right+1]
}

func findClosestElementIdx(arr []int, x int) int {
	if x <= arr[0] {
		return 0
	}
	if x >= arr[len(arr)-1] {
		return len(arr) - 1
	}
	bestIdx := 0
	lo := 0
	hi := len(arr) - 1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		// if arr[mid] == x {
		// 	return mid
		// } else
		if x <= arr[mid] {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
		if abs(arr[mid]-x) < abs(arr[bestIdx]-x) {
			bestIdx = mid

		} else if abs(arr[mid]-x) == abs(arr[bestIdx]-x) {
			if mid < bestIdx {
				bestIdx = mid
			}
		}
	}
	return bestIdx
}

func abs(x int) int {
	if x < 0 {
		x *= -1
	}
	return x
}
