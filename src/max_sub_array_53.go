package main

import "math"

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	res := nums[0]
	sum := 0
	for _, num := range nums {
		if sum > 0 {
			sum += num
		} else {
			sum = num
		}
		res = (int)(math.Max(float64(res), float64(sum)))
	}
	return res
}