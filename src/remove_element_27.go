package main

import "fmt"

func main() {
	var nums = []int{1}
	fmt.Println(removeElement(nums, 5))
	fmt.Println(nums)
}

func removeElement(nums []int, val int) int {
	var pos = 0
	for _, num := range nums {
		if num != val {
			nums[pos] = num
			pos ++
		}
	}
	return pos
}
