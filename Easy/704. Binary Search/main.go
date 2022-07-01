package _04__Binary_Search

func search(nums []int, target int) int {
	var left, pivot int
	right := len(nums) - 1
	for left <= right {
		pivot = left + (right-left)/2
		if nums[pivot] == target {
			return pivot
		}
		if nums[pivot] > target {
			right = pivot - 1
		} else {
			left = pivot + 1
		}
	}
	return -1
}
