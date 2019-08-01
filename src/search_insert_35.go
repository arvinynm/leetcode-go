package main

func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	for i, num := range nums {
		if num == target {
			return i
		}
		if num > target {
			return i - 1
		}
	}
	return len(nums)
}
