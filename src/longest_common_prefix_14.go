package main

import "fmt"

func main() {
	fmt.Println()
}

func longestCommonPrefix(strs []string) string {
	ret := []byte{}
	if len(strs) == 0 {
		return string(ret)
	}
	min := len(strs[0])
	for _, str := range strs {
		if len(str) < min {
			min = len(str)
		}
	}
flag:
	for i := 0; i < min; i++ {
		f := strs[0][i]
		for _, str := range strs {
			if str[i] != f {
				break flag
			}
		}
		ret = append(ret, f)
	}
	return string(ret)
}
