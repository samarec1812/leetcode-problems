package search_insert_position

/*
link problem: https://leetcode.com/problems/search-insert-position/description
*/

func SearchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		middle := (right + left) / 2
		if nums[middle] < target {
			left = middle + 1
		} else if nums[middle] > target {
			right = middle - 1
		} else {
			return middle
		}
	}
	return left
}
