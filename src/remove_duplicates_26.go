package main

import "fmt"

func main (){
	fmt.Println()
	nums := []int {1,2,3,4,4,5,5,6,6,6,7}
	fmt.Println(removeDuplicates(nums))
	fmt.Println(nums)
}

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var pos int
	for i, _ := range nums {
		if nums[i] != nums[pos] {
			pos += 1
			nums[pos] = nums[i]
		}
	}
	return pos + 1
}