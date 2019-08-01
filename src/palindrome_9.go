package main

func main() {
	isPalindrome(10)
}
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	array := make([]int, 0)
	for {
		b := x / 10
		a := x % 10
		array = append(array, a)
		if b == 0 {
			break

		}
		x = b
	}
	for i, _ := range array {
		if array[i] != array[len(array)-i-1] {
			return false
		}
	}
	return true
}