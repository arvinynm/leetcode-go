package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strStr("abcsdasd", "sd"))
}

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	if !strings.Contains(haystack, needle) {
		return -1
	}
	for i, _ := range haystack {

		for j := 0; j < len(needle); j++ {
			if haystack[i + j] != needle[j] {
				break
			}
			if j == len(needle) - 1 {
				return i
			}
		}
	}
	return -1
}